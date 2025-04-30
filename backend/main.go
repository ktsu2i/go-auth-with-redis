package main

import (
	"backend/db"
	"backend/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		e.Logger.Fatal(err)
	}
	handler.DB = db

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	route(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
