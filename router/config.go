package router

import (
	"github.com/wasabi315/todo-server/repository"
)

type Config struct {
	Repo repository.Repository
}
