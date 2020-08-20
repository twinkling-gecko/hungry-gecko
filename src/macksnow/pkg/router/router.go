package router

import (
	"github.com/labstack/echo/v4"

	"macksnow/pkg/repository"
)

var repo repository.Repository

func Init(e *echo.Echo, repository repository.Repository) {
	repo = repository

	itemsRouter(e)
}
