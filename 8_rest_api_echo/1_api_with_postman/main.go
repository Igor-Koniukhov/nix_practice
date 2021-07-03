package main

import (
	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	"Nix_practice/8_rest_api_echo/1_api_with_postman/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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


func HandleRoutes(){
	e := echo.New()
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