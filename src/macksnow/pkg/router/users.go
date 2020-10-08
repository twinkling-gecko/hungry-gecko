package router

import (
	"github.com/labstack/echo/v4"
)

type receiveUser struct {
	Nickname   string `json:"nickname" validate: "required"`
	ScreenName string `json:"screen_name" validate: "required"`
	Email      string `json:"email" validate: "required"`
	Password   string `json:"password" validate: "required"`
}

// @summary Register user info
// @produce json
// @param users body receiveUser true "User data"
// @success 200 {object} model.User
// @failure 400 {object} errorResponse
// @router /api/v1/users [post]
func usersCreateRouter(e *echo.Echo) {
	e.POST("/v1/users", func(c echo.Context) error {
		// リクエストバインド
		// リクエストのバリデート
		// パスワードハッシュ
		// Userモデルを起こす
		// Userモデルバリデート
		// CreateUser
		// return
		return nil
	})
}
