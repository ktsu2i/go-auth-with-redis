package main

import (
	"backend/db"
	"backend/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		e.Logger.Fatal(err)
	}
	handler.DB = db

	rc := redis.NewClient(&redis.Options{
		Addr: "auth_redis:6379",
		DB:   0,
	})
	handler.RC = rc

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	route(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
