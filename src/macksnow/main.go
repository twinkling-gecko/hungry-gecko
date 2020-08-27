package main

import (
	"net/http"

	echoSwagger "github.com/ken-aio/echo-swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "macksnow/docs"
	"macksnow/pkg/repository"
	"macksnow/pkg/router"
)

const PORT = ":4000"

func main() {
	e := echo.New()

	repo, err := repository.New()
	if err != nil {
		panic(err)
	}

	// TODO: DB切断すらできなかったときのエラーハンドリング
	defer repo.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/docs/*", echoSwagger.WrapHandler) // swagger documents
	router.Init(e, repo)

	e.Logger.Fatal(e.Start(PORT))
}

func hello(c echo.Context) error {
	type response struct {
		Message string `json:"message"`
	}

	return c.JSON(http.StatusOK, &response{Message: "I'm fine."})
}
