package router

import (
	"net/http"
	"os"

	"todo/app/application/interface/database"
	"todo/app/application/schema"
	"todo/app/application/usecase"
	"todo/app/presentation/controller"
	appmiddleware "todo/app/presentation/middleware"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func SetUp(e *echo.Echo, logger *zap.Logger, databaseHandller database.IDatabaseHandller) error {
	authUsecase := usecase.NewAuthUsecase(databaseHandller)
	taskUsecase := usecase.NewTaskUsecase(databaseHandller)
	priorityUsecase := usecase.NewPriorityUsecase(databaseHandller)

	csrfController := controller.NewCsrfController()
	authController := controller.NewAuthController(authUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	priorityController := controller.NewPriorityController(priorityUsecase)

	// apiグループ
	api := e.Group("/api")

	// CORSミドルウェアの設定
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFミドルウェアの設定
	api.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("MY_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieMaxAge:   60,
	}))

	// リクエストミドルウェアの設定
	api.Use(appmiddleware.Request(logger))

	// default
	api.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, schema.DefaultResponseModel{
			Message: "Hello, world!",
		})
	})

	// CSRF関連
	api.GET("/csrf", csrfController.GetCsrfToken)

	// 認証関連
	api.POST("/auth/signup", authController.SignUp)
	api.POST("/auth/signin", authController.SignIn)
	api.GET("/auth/signout", authController.SignOut)

	// v1グループ
	v1 := api.Group("/v1")

	// JWTミドルウェアの設定
	v1.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_KEY")),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, schema.DefaultResponseModel{
				Message: err.Error(),
			})
		},
	}))

	// タスク関連
	v1.POST("/task", taskController.CreateTask)
	v1.PUT("/task", taskController.UpdateTask)
	v1.GET("task/:id", taskController.GetTask)
	v1.DELETE("task/:id", taskController.DeleteTask)
	v1.GET("task/done/:id", taskController.DeleteTask)
	v1.GET("/tasks", taskController.GetAllTask)

	// 優先度関連
	v1.GET("/priorities", priorityController.GetAllPriority)

	return nil
}
