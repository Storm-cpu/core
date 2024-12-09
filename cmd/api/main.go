package main

import (
	"money-note/config"
	"money-note/pkg/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadConfig()

	e := server.New(&server.Config{
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.Start(e)
}
