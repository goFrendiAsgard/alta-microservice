package main

import "github.com/labstack/echo"

func SetRoute(config Config, e *echo.Echo) {
	if config.Mode == "" || config.Mode == "reader" {
		e.GET("news", CreateGetNewsController(config))
	}
	if config.Mode == "" || config.Mode == "writer" {
		e.POST("news", CreatePostNewsController(config))
	}
	if config.Mode == "gateway" {
		e.GET("news", CreateGatewayGetNewsController(config))
		e.POST("news", CreateGatewayPostNewsController(config))
	}
}
