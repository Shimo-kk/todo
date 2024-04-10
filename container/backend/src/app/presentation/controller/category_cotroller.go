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

type categoryController struct {
	usecase usecase.ICategoryUsecase
}

// カテゴリコントローラーの作成
func NewCategoryController(usecase usecase.ICategoryUsecase) *categoryController {
	return &categoryController{usecase: usecase}
}

// カテゴリの作成
func (c *categoryController) CreateCategory(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.CategoryCreateModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.CreateCategory(userId, requestBody); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("カテゴリを作成しました。"))
}

// カテゴリの更新
func (c *categoryController) UpdateCategory(context echo.Context) error {
	// リクエストボディをバインド
	requestBody := schema.CategoryUpdateModel{}
	if err := context.Bind(&requestBody); err != nil {
		return core.NewError(core.BadRequestError, err.Error())
	}

	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケースを実行
	if err := c.usecase.UpdateCategory(userId, requestBody); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("カテゴリを更新しました。"))
}

// カテゴリの取得
func (c *categoryController) GetCategory(context echo.Context) error {
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
	responseBody, err := c.usecase.GetCategory(userId, id)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}

// カテゴリの削除
func (c *categoryController) DeleteCategory(context echo.Context) error {
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
	if err := c.usecase.DeleteCategory(userId, id); err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responce.NewDefaultRespoce("カテゴリを削除しました。"))
}

// タスクを全件取得
func (c *categoryController) GetAllCategory(context echo.Context) error {
	// JWTトークンからユーザーIDを取得
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	// ユースケース実行
	responseBody, err := c.usecase.GetAllCategory(userId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, responseBody)
}
