package routes

import (
	"github.com/goravel/framework/facades"
)

func Api() {
	facades.Route().Static("static", "./public")
	facades.Route().StaticFile("favicon.ico", "./public/favicon.ico")
}
