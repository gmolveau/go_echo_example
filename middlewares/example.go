package middlewares

import "github.com/labstack/echo"

func Example() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// do things here
			return next(c)
		}
	}
}