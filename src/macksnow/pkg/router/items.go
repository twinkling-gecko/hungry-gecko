package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// TODO: ダミーAPIの定義用なのでいずれ消す
type item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Summary   string    `json:"summary"`
	URI       string    `json:"uri"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type indexResponse struct {
	Items []*item `json:"items"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// TODO: ダミーAPI用のダミーデータなのでいずれ消す
var sampleItem = &item{
	ID:        0,
	Name:      "TestItem",
	Summary:   "TestSummary",
	URI:       "https://www.amazon.co.jp",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func itemsRouter(e *echo.Echo) {
	itemsIndexRouter(e)
	itemsShowRouter(e)
}

// @summary Get itemslist
// @produce json
// @success 200 {object} indexResponse
// @router /api/v1/items [get]
func itemsIndexRouter(e *echo.Echo) {
	e.GET("/v1/items", func(c echo.Context) error {
		// TODO: 本実装

		items := []*item{sampleItem}
		res := &indexResponse{Items: items}
		return c.JSON(http.StatusOK, res)
	})
}

func itemsShowRouter(e *echo.Echo) {
	e.GET("/v1/items/:id", func(c echo.Context) error {
		// TODO: 本実装

		req := c.Param("id")

		id, err := strconv.Atoi(req)
		if err != nil {
			res := &errorResponse{Message: req + " is not integer."}
			return c.JSON(http.StatusBadRequest, res)
		}

		sampleItem := &item{
			ID:        id,
			Name:      "TestItem",
			Summary:   "TestSummary",
			URI:       "https://www.amazon.co.jp",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		return c.JSON(http.StatusOK, sampleItem)
	})
}
