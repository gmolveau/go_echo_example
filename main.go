package main

import (
	"github.com/gmolveau/go_echo_example/controllers"
	"github.com/gmolveau/go_echo_example/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middlewares.Example())

	api := e.Group("/api")
	{
		r_example := api.Group("/examples")
		{
			r_example.GET("", controllers.ListExample)
			r_example.POST("", controllers.CreateExample)
			r_example.GET("/:id", controllers.GetExample)
			r_example.PUT("/:id", controllers.UpdateExample)
			r_example.DELETE("/:id", controllers.DeleteExample)
		}
	}
	e.Logger.Fatal(e.Start(":1337"))
}
