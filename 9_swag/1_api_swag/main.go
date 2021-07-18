package main

import (
	"github.com/Igor-Koniukhov/api_with_swag/dbase"
	"github.com/Igor-Koniukhov/api_with_swag/internal/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"

	"database/sql"
	"fmt"
	_ "github.com/Igor-Koniukhov/api_with_swag/docs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)
// @title Swagger REST API
// @version 1.0
// @description This is a sample of REST API
// @host localhost:8080
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /


var err error

func main() {
	dbase.Db, err = sql.Open(dbase.Driver, dbase.DbSN)
	if err == nil {
		fmt.Println("Success! DB connection on port:", dbase.PortDB)
	} else {
		fmt.Println(err)
	}
	HandleRoutes()
}

func HandleRoutes() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowOrigins: []string{"*"},
	}))

	// Routes
	e.GET("/post", handlers.GetAllUsers)
	e.POST("/post", handlers.CreatePost)
	e.GET("/post/:id", handlers.GetPost)
	e.PUT("/post/:id", handlers.UpdatePost)
	e.DELETE("/post/:id", handlers.DeletePost)

	e.POST("/comment", handlers.CreateComment)
	e.GET("/comment/:id", handlers.GetComment)
	e.PUT("/comment/:id", handlers.UpdateComment)
	e.DELETE("/comment/:id", handlers.DeleteComment)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
