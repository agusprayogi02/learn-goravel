package models

import (
	"github.com/goravel/framework/database/orm"
)

type Todo struct {
	orm.Model
	Title   string
	Content string
	orm.SoftDeletes
}
