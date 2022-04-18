package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gba-3/sample-todo/domain/entity"
	"github.com/gba-3/sample-todo/usecase"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	uu usecase.UserUsecase
}

type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	Signup(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
	GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error)
}

func NewUserHandler(uu usecase.UserUsecase) UserHandler {
	return &userHandler{uu}
}

func (uh *userHandler) Login(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	user := entity.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if user.Email == "" {
		return http.StatusBadRequest, nil, errors.New("User struct error: email is empty.")
	}
	if user.Password == "" {
		return http.StatusBadRequest, nil, errors.New("User struct error: password is empty.")
	}
	u, err := uh.uu.Find(ctx, user.Email)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(u.ID),
		ExpiresAt: time.Now().In(jst).Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	res := map[string]string{
		"token": token,
	}
	return http.StatusOK, res, nil
}

func (uh *userHandler) Signup(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	user := entity.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if user.Name == "" {
		return http.StatusBadRequest, nil, errors.New("User struct error: name is empty.")
	}
	if user.Email == "" {
		return http.StatusBadRequest, nil, errors.New("User struct error: email is empty.")
	}
	if user.Password == "" {
		return http.StatusBadRequest, nil, errors.New("User struct error: password is empty.")
	}

	// GenerateFromPasswordでパスワードをハッシュ化
	// 第2引数はコスト
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if err := uh.uu.Create(ctx, user.Name, user.Email, string(hashed)); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, "success", nil
}

func (uh *userHandler) GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ctx := context.Background()
	users, err := uh.uu.GetAll(ctx)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, users, nil
}
