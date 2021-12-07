package infrastructure

import (
	"github.com/taiki1288/Go-CleanArchitecture-APIServer/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	userController := controllers.NewUserController(NewSqlHandler())

	// middleware.Loggerはアクセスログのようなリクエスト単位のログを出力する。
	e.Use(middleware.Logger())
	// middleware.Recoverはアプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする。
	e.Use(middleware.Recover())

	e.POST("/users", func(c echo.Context) error { return userController.CreateUser(c) })
	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.GetUser(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.UpdateUser(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.DeleteUser(c) })

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}