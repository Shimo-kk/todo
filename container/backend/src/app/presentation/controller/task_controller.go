package controller

import (
	"net/http"
	"strconv"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
	"todo/app/presentation/responce"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type taskController struct {
	usecase usecase.ITaskUsecase
}

// タスクコントローラーの作成
func NewTaskController(usecase usecase.ITaskUsecase) *taskController {
	return &taskController{usecase: usecase}
}

// タスクの作成
func (c *taskController) CreateTask(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.TaskCreateModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.CreateTask(userId, requestBody); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("タスクを作成しました。"))
}

// タスクの更新
func (c *taskController) UpdateTask(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.TaskUpdateModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.UpdateTask(userId, requestBody); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("タスクを更新しました。"))
}

// タスクの取得
func (c *taskController) GetTask(context echo.Context) error {
	// パラメータを取得
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	responseBody, err := c.usecase.GetTask(userId, id)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}

// タスクの削除
func (c *taskController) DeleteTask(context echo.Context) error {
	// パラメータを取得
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.DeleteTask(userId, id); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("タスクを削除しました。"))
}

// タスクの完了
func (c *taskController) DoneTask(context echo.Context) error {
	// パラメータを取得
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.DoneTask(userId, id); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("タスクを完了しました。"))
}

// タスクを全件取得
func (c *taskController) GetAllTask(context echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケース実行
	responseBody, err := c.usecase.GetAllTask(userId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}
