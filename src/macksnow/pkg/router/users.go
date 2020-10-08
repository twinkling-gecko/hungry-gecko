package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type receiveUser struct {
	Nickname   string `json:"nickname" validate: "required,min=4,max=15"`
	ScreenName string `json:"screen_name" validate: "required,alphanum,min=4,max=15"`
	Email      string `json:"email" validate: "required,email"`
	Password   string `json:"password" validate: "required,containsany=!@#$%^&*()-=_+[]{}\|;':",.<>/?,min=8,max=20"`
}

// @summary Register user info
// @produce json
// @param users body receiveUser true "User data"
// @success 200 {object} model.User
// @failure 400 {object} errorResponse
// @router /api/v1/users [post]
func usersCreateRouter(e *echo.Echo) {
	e.POST("/v1/users", func(c echo.Context) error {

		receivedUser := new(receiveUser)

		// リクエストのバインド
		if err := c.Bind(receivedUser); err != nil {
			res := &errorResponse{Message: "Invalid parameters."}
			return c.JSON(http.StatusBadRequest, res)
		}

		// リクエストのバリデート
		if err := c.Validate(receivedUser); err != nil {
			// TODO: 不正パラメータの原因を通知する
			res := &errorResponse{Message: "The requested parameter is invalid." + err.Error()}
			return c.JSON(http.StatusBadRequest, res)
		}

		// パスワードハッシュの生成
		// confirmation_tokenの生成
		// Userモデルを起こす
		// Userモデルバリデート
		// CreateUser
		// return
		return nil
	})
}
