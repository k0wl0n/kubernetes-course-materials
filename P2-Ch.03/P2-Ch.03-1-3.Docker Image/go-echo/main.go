package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			name = "World"
		}
		return c.String(http.StatusOK, "Hello, "+name+"!")
	})

	e.GET("/v1", func(c echo.Context) error {
		age := c.QueryParam("age")
		if age == "" {
			age = "10"
		}
		return c.String(http.StatusOK, "Hello, your age "+age+"!")
	})

	e.Start(":8080")
}
