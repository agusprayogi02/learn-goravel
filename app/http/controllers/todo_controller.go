package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type TodoController struct {
	//Dependent services
}

func NewTodoController() *TodoController {
	return &TodoController{
		//Inject services
	}
}

func (r *TodoController) Index(ctx http.Context) http.Response {
	return nil
}	

func (r *TodoController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *TodoController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *TodoController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *TodoController) Destroy(ctx http.Context) http.Response {
	return nil
}
