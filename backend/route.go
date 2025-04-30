package main

import (
	"backend/handler"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	api := e.Group("/api")

	api.POST("/signup", handler.SignUp)
}
