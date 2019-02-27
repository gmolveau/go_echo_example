package controllers

import "github.com/labstack/echo"

func ListExample(c echo.Context) error {
	// do things here
	return c.JSON(200, nil)
}

func GetExample(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debug(id)
	// do things here
	return c.JSON(200, nil)
}

func CreateExample(c echo.Context) error {
	// c.Bind(your_struct)
	// do things here
	return c.JSON(200, nil)
}

func UpdateExample(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debug(id)
	// c.Bind(your_struct)
	// do things here
	return c.JSON(200, nil)
}

func DeleteExample(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Debug(id)
	// do things
	return c.JSON(200, nil)
}