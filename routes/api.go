package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Static("static", "./public")
	facades.Route().StaticFile("favicon.ico", "./public/favicon.ico")
}
