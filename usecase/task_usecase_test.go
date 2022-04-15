package usecase

import (
	"context"
	"testing"

	"github.com/gba-3/sample-todo/domain/entity"
)

type mockTaskRepository struct {
	changeStatusFunc func() error
	addTaskFunc      func() error
	getAllFunc       func() (entity.Tasks, error)
}

func (tr *mockTaskRepository) ChangeStatus(ctx context.Context, id int, status bool) error {
	return tr.changeStatusFunc()
}

func (tr *mockTaskRepository) AddTask(ctx context.Context, title string, description string, date string) error {
	return tr.addTaskFunc()
}

func (tr *mockTaskRepository) GetAll(ctx context.Context) (entity.Tasks, error) {
	return tr.getAllFunc()
}
func TestAddTask(t *testing.T) {
	testCases := []struct {
		tr          *mockTaskRepository
		title       string
		description string
		date        string
		expected    error
	}{
		{
			&mockTaskRepository{
				addTaskFunc: func() error {
					return nil
				},
			},
			"test title",
			"test description",
			"2022-04-11 15:00:00",
			nil,
		},
	}

	ctx := context.Background()
	for _, testCase := range testCases {
		tu := NewTaskUsecase(testCase.tr)
		if err := tu.AddTask(ctx, testCase.title, testCase.description, testCase.date); err != nil {
			t.Errorf("unexpected result. actual: %t, expected: %t", err, testCase.expected)
		}
	}
}
