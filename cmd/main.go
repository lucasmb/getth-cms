package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucasmb/getth-cms/db"
	"github.com/lucasmb/getth-cms/handler"
	"github.com/lucasmb/getth-cms/service"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DB_NAME string = "./db/echo.db"

func main() {

	app := echo.New()

	//**** Middlewares ****/
	app.Use(middleware.Logger())

	// app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:4000", "https://localhost:2015", "https://sabado.dev", "https://sabado.dev"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	db := db.InitDb(DB_NAME)
	// store, err := db.NewStore(DB_NAME)
	// if err != nil {
	// 	app.Logger.Fatalf("failed to create Db store: %s", err)
	// }


	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to GETTH-CMS!")
	})

	userService := service.NewUserService(db)
	handler.AuthRegisterHandler(app, handler.NewAuthHandler(userService))
	handler.UserRegisterHandler(app, handler.NewUserHandler(userService))

	postService := service.NewPostService(db)
	handler.PostRegisterHandler(app, handler.NewPostHandler(postService))

	//Dashboard
	handler.DashboardRegisterHandler(app, handler.NewDashboardHandler(userService, postService))

	//serve assets css
	app.Static("/assets", "assets")

	app.Logger.Fatal(app.Start(":4000"))

}
