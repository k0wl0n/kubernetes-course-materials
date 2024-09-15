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
			name = "echo"
		}
		return c.String(http.StatusOK, "Hello, "+name+"! on v0.2")
	})

	e.GET("/v1", func(c echo.Context) error {
		age := c.QueryParam("age")
		if age == "" {
			age = "15"
		}
		return c.String(http.StatusOK, "Hello, your age "+age+"!")
	})

	e.Start(":8080")
}
