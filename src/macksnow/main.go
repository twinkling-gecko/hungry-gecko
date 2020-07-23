package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"macksnow/pkg/router"
)

const PORT = ":4000"

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	router.Init(e)

	e.Logger.Fatal(e.Start(PORT))
}

func hello(c echo.Context) error {
	type response struct {
		Message string `json:"message"`
	}

	return c.JSON(http.StatusOK, &response{Message: "I'm fine."})
}
