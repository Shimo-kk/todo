package controller

import (
	"net/http"
	"os"
	"time"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
	"todo/app/presentation/responce"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type authController struct {
	usecase usecase.IAuthUsecase
}

// 認証コントローラーの作成
func NewAuthController(usecase usecase.IAuthUsecase) *authController {
	return &authController{usecase: usecase}
}

// サインアップ
func (c *authController) SignUp(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.SignUpModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// ユースケースを実行
	if err := c.usecase.SignUp(requestBody); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("サインアップしました。"))
}

// サインイン
func (c *authController) SignIn(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.SignInModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// ユースケースを実行
	userId, err := c.usecase.SignIn(requestBody)
	if err != nil {
		return err
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return core.NewError(core.SystemError, err.Error())
	}

	// CookieにJWTトークンを設定
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenStr
	cookie.Expires = (time.Now().Add(24 * time.Hour))
	cookie.Path = "/"
	cookie.Domain = os.Getenv("MY_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	context.SetCookie(cookie)

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("サインインしました。"))
}

// サインアウト
func (ac *authController) SignOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("MY_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, responce.NewDefaultRespoce("サインアウトしました。"))
}
