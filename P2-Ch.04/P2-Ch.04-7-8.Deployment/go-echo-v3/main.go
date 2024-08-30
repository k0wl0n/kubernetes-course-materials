package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			name = "World"
		}
		hostname, _ := os.Hostname()
		env := os.Getenv("ENV")
		return c.String(http.StatusOK, "Hello, "+name+"! using v3 on "+hostname+" in "+env+" environment")
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
