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
		env := os.Getenv("MISSING_ENV")
		if env == "" {
			// Simulate a critical error by exiting the application
			os.Exit(1)
		}
		return c.String(http.StatusOK, "Hello, "+name+"! using v4 on "+hostname+" in "+env+" environment")
	})

	e.GET("/health", func(c echo.Context) error {
		hostname, _ := os.Hostname()
		return c.String(http.StatusOK, hostname)
	})

	e.Start(":8080")
}
