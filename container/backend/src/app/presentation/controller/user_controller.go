package controller

import (
	"net/http"
	"todo/app/application/interface/usecase"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type userController struct {
	usecase usecase.IUserUsecase
}

// ユーザーコントローラーの作成
func NewUserController(usecase usecase.IUserUsecase) *userController {
	return &userController{usecase: usecase}
}

// ユーザー取得
func (c *userController) GetUser(context echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	responseBody, err := c.usecase.GetUser(userId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}
