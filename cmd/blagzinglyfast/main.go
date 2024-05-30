package main

import (
	"html/template"
	"io"

	"github.com/Suanlianmang/blagzignlyfast/pkg/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


func newTemplates(pattern string) *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob(pattern)),
	}
}

func rendererMiddleware(pattern string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
            renderer := newTemplates(pattern)
			c.Echo().Renderer = renderer
			return next(c)
		}
	}
}


func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Static("/static", "assets")
    
    count := &pages.Count{ Count: 0 }

    base := e.Group("/")
    // Middleware to pass the count to handlers
    base.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Set("count", count)
            return next(c)
        }
    })
    base.Use(rendererMiddleware("views/test.html"))

    base.GET("", pages.Index)
    base.POST("count", pages.Increment)

   
    e.Logger.Fatal(e.Start(":8000"))

}
