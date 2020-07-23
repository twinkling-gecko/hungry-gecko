package router

import (
	"net/http"
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
