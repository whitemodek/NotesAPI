package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)


func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}


func createNote(c echo.Context) error {
	return c.String(http.StatusCreated, "Here will be note creating")
}


func showNote(c echo.Context) error {
	id := c.QueryParam("id")

	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "The ID parameter is required")
	}

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The ID must be a number")
	}

	if idNum < 1 {
		return echo.NewHTTPError(http.StatusNotFound, "Note is not found: The ID must be greater than 0")
	}

	return c.String(http.StatusOK, "Displaying a note with an ID "+id)
}