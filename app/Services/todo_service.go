package services

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Create(req requests.TodoRequest) error {
	todo := models.Todo{
		Title:   req.Title,
		Content: req.Content,
	}
	tx, err := facades.Orm().Query().Begin()
	if err != nil {
		return err
	}
	if err := tx.Create(&todo); err != nil {
		return tx.Rollback()
	} else {
		return tx.Commit()
	}
}
