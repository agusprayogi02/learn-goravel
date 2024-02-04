package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type TodoRequest struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

func (r *TodoRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *TodoRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TodoRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TodoRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *TodoRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
