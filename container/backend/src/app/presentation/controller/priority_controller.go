package controller

import (
	"net/http"
	"todo/app/application/interface/usecase"

	"github.com/labstack/echo/v4"
)

type priorityController struct {
	usecase usecase.IPriorityUsecase
}

// 優先度コントローラーの作成
func NewPriorityController(usecase usecase.IPriorityUsecase) *priorityController {
	return &priorityController{usecase: usecase}
}

// 優先度を全件取得
func (c *priorityController) GetAllPriority(context echo.Context) error {
	// ユースケース実行
	responseBody, err := c.usecase.GetAllPriority()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}
