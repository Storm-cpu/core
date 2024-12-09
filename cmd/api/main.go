package main

import (
	"net/http"

	"github.com/Storm-cpu/core/config"
	"github.com/Storm-cpu/core/pkg/server"

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
