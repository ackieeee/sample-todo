package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func WithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// ctx := context.WithValue(r.Context(), "token", tokenString)
		_, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		// next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func CheckToken(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Authorization:Bearer xxxx からトークンを抽出
// 		tokenString := r.Header.Get("Authorization")
// 		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 		// tokenの認証
// 		token, err := verifyToken(tokenString)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		// ペイロード読み出し
// 		claims := token.Claims.(jwt.MapClaims)
// 		next.ServeHTTP(w, r)
// 	})
// }
