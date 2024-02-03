package controllers

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": strings.Join(strings.Split(ctx.Request().Route("id"), "-"), " "),
	})
}
