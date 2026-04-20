package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", home)
	e.GET("/note", showNote)
	e.POST("/note/create", createNote)

	e.Logger.Fatal(e.Start(":4000"))
}	