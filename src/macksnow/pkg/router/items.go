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
// @failure 500 {object} errorResponse
// @router /api/v1/items [get]
func itemsIndexRouter(e *echo.Echo) {
	e.GET("/v1/items", func(c echo.Context) error {
		items, err := repo.AllItems()
		if err != nil {
			res := &errorResponse{Message: err.Error()}
			return c.JSON(http.StatusInternalServerError, res)
		}

		res := &indexResponse{Items: items}
		return c.JSON(http.StatusOK, res)
	})
}

// @summary Get item info
// @produce json
// @param id path integer true "Item ID"
// @success 200 {object} model.Item
// @failure 400 {object} errorResponse
// @failure 404 {object} errorResponse
// @router /api/v1/items/{id} [get]
func itemsShowRouter(e *echo.Echo) {
	e.GET("/v1/items/:id", func(c echo.Context) error {
		req := c.Param("id")

		id, err := strconv.Atoi(req)
		if err != nil {
			res := &errorResponse{Message: req + " is not integer."}
			return c.JSON(http.StatusBadRequest, res)
		}

		item, err := repo.FindItem(id)
		if err != nil {
			// TODO: 内部エラーとレコード無しのハンドリングをする
			res := &errorResponse{Message: err.Error()}
			return c.JSON(http.StatusNotFound, res)
		}

		return c.JSON(http.StatusOK, item)
	})
}

// @summary Register item info
// @produce json
// @param items body receiveItem true "Item data"
// @success 200 {object} model.Item
// @failure 400 {object} errorResponse
// @router /api/v1/items [post]
func itemsCreateRouter(e *echo.Echo) {
	e.POST("/v1/items", func(c echo.Context) error {

		// echoにvalidatorを登録
		e.Validator = &CustomValidator{validator: validator.New()}

		newItem := new(model.Item)

		if err := c.Bind(newItem); err != nil {
			res := &errorResponse{Message: "Invalid parameters."}
			return c.JSON(http.StatusBadRequest, res)
		}

		if err := c.Validate(newItem); err != nil {
			res := &errorResponse{Message: "Required parameters is empty. " + err.Error()}
			return c.JSON(http.StatusBadRequest, res)
		}

		item, err := repo.CreateItem(newItem.Name, newItem.Summary, newItem.URI)
		if err != nil {
			res := &errorResponse{Message: err.Error()}
			return c.JSON(http.StatusInternalServerError, res)
		}

		return c.JSON(http.StatusOK, item)
	})
}
