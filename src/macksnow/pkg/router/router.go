package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	itemsRouter(e)
}

// TODO: いずれ消す
type item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Summary   string    `json:"summary"`
	URI       string    `json:"uri"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func itemsRouter(e *echo.Echo) {
	e.GET("/v1/items", func(c echo.Context) error {
		type response struct {
			Items []*item `json:"items"`
		}

		items := []*item{
			{
				ID:        0,
				Name:      "TestItem",
				Summary:   "TestSummary",
				URI:       "https://www.amazon.co.jp",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		res := &response{Items: items}
		return c.JSON(http.StatusOK, res)
	})
}
