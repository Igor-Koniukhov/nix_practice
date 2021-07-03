package main

import (
	"Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"
	"Nix_practice/8_rest_api_echo/2_api_with_templ/internal/handlers"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
)

var (
	err  error
	port = ":8000"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	dbase.Db, err = sql.Open(dbase.Driver, dbase.DbSN)
	if err == nil {
		fmt.Println("Success! DB connection on port:", dbase.PortDB)
	} else {
		fmt.Println(err)
	}

	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowOrigins: []string{"*"},
	}))

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", handlers.Home)
	e.GET("/users/post-json/", handlers.GetPostJSON)
	e.GET("/users/post-xml/", handlers.GetPostXML)
	e.GET("/users/post/get/", handlers.GetForUpdatePost)
	e.POST("/users/post/update/", handlers.UpdatePost)

	e.GET("/users/", handlers.DeletePost)

	e.GET("/users/post/comment/", handlers.WriteComment)
	e.POST("/users/post/create/comment/", handlers.CreateComment)

	e.GET("/users/post/write/", handlers.WritePost)
	e.POST("/users/post/create/", handlers.CreatePost)

	e.GET("/users/comment-json/:id", handlers.GetCommentJSON)
	e.GET("/users/comment-xml/:id", handlers.GetCommentXML)

	e.GET("/users/comment/get/", handlers.GetForUpdateComment)
	e.POST("/users/comment/get/", handlers.UpdateComment)
	e.GET("/users/comment/delete/", handlers.DeleteComment)

	e.Logger.Fatal(e.Start(port))
}


