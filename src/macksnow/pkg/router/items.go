package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"macksnow/pkg/model"
)

type receiveItem struct {
	Name    string `json:"name" validate:"required"`
	Summary string `json:"summary" validate:"required"`
	URI     string `json:"uri" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

type indexResponse struct {
	Items []*model.Item `json:"items"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// TODO: ダミーAPI用のダミーデータなのでいずれ消す
var sampleItem = &model.Item{
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
	itemsCreateRouter(e)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @summary Get itemslist
// @produce json
// @success 200 {object} indexResponse
// @router /api/v1/items [get]
func itemsIndexRouter(e *echo.Echo) {
	e.GET("/v1/items", func(c echo.Context) error {
		// TODO: 本実装

		items := []*model.Item{sampleItem}
		res := &indexResponse{Items: items}
		return c.JSON(http.StatusOK, res)
	})
}

// @summary Get item info
// @produce json
// @param id query integer true "Item ID"
// @success 200 {object} item
// @failure 400 {object} errorResponse
// @router /api/v1/items/{id} [get]
func itemsShowRouter(e *echo.Echo) {
	e.GET("/v1/items/:id", func(c echo.Context) error {
		// TODO: 本実装

		req := c.Param("id")

		id, err := strconv.Atoi(req)
		if err != nil {
			res := &errorResponse{Message: req + " is not integer."}
			return c.JSON(http.StatusBadRequest, res)
		}

		sampleItem := &model.Item{
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

// @summary Register item info
// @produce json
// @param items body receiveItem true "Item data"
// @success 200 {object} item
// @failure 400 {object} errorResponse
// @router /api/v1/items [post]
func itemsCreateRouter(e *echo.Echo) {
	e.POST("/v1/items", func(c echo.Context) error {
		// TODO: 本実装

		// echoにvalidatorを登録
		e.Validator = &CustomValidator{validator: validator.New()}

		sampleItem := new(model.Item)

		if err := c.Bind(sampleItem); err != nil {
			res := &errorResponse{Message: "Invalid parameters."}
			return c.JSON(http.StatusBadRequest, res)
		}

		if err := c.Validate(sampleItem); err != nil {
			res := &errorResponse{Message: "Required parameters is empty. " + err.Error()}
			return c.JSON(http.StatusBadRequest, res)
		}

		sampleItem.ID = 99
		sampleItem.CreatedAt = time.Now()
		sampleItem.UpdatedAt = time.Now()
		return c.JSON(http.StatusOK, sampleItem)
	})
}
