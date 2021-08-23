package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := NewConfig()
	e := echo.New()
	e.Use(middleware.Logger())
	SetRoute(config, e)
	// serve
	port := config.GetPort()
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
