package controllers

import (
	services "goravel/app/Services"
	"goravel/app/http/requests"

	"github.com/goravel/framework/contracts/http"
)

type TodoController struct {
	// Dependent services
	service services.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{
		// Inject services
		service: *services.NewTodoService(),
	}
}

func (r *TodoController) Index(ctx http.Context) http.Response {
	return ctx.Response().View().Make("todo/index.tmpl", map[string]any{
		"halo": "Semua orang bisa makan",
	})
}

func (r *TodoController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *TodoController) Store(ctx http.Context) http.Response {
	var todo *requests.TodoRequest
	if err := ctx.Request().Bind(&todo); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
		})
	}
	if err := r.service.Create(*todo); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err.Error(),
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"message": "Berhasil",
	})
}

func (r *TodoController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *TodoController) Destroy(ctx http.Context) http.Response {
	return nil
}
