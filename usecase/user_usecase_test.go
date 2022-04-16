package usecase

import (
	"context"
	"testing"

	"github.com/gba-3/sample-todo/domain/entity"
)

type mockUserRepository struct {
	getAllFunc func() (entity.Users, error)
}

func (ur *mockUserRepository) GetAll(ctx context.Context) (entity.Users, error) {
	return ur.getAllFunc()
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		ur            *mockUserRepository
		expectedID    int
		expectedError error
	}{
		{
			&mockUserRepository{
				func() (entity.Users, error) {
					users := entity.Users{
						{
							ID: 1,
						},
					}
					return users, nil
				},
			},
			1,
			nil,
		},
	}

	ctx := context.Background()
	for _, testCase := range testCases {
		tu := NewUserUsecase(testCase.ur)
		tasks, err := tu.GetAll(ctx)
		if err != nil {
			t.Errorf("unexpected result. expected: %t, actual: %t", testCase.expectedError, err)
		}
		if len(tasks) == 0 {
			t.Errorf("task is empty.")
		}
	}
}