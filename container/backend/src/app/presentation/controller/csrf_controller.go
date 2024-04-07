package controller

import (
	"net/http"
	"todo/app/application/schema"

	"github.com/labstack/echo/v4"
)

type csrfController struct {
}

// 認証コントローラーの作成
func NewCsrfController() *csrfController {
	return &csrfController{}
}

// CSRFトークンの取得
func (c *csrfController) GetCsrfToken(context echo.Context) error {
	token := context.Get("csrf").(string)

	resonseBody := schema.CSRFModel{
		Csrf: token,
	}
	return context.JSON(http.StatusOK, resonseBody)
}
