package main

import (
	"github.com/gmolveau/go_echo_example/controllers"
	"github.com/gmolveau/go_echo_example/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middlewares.Example())

	t.Run("ListExample", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/examples")
		// assertions
		if assert.NoError(t, controllers.ListExample(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	// add more unit tests here
}