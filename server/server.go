package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	log.Println(c.Request().URL)
	if c.Request().URL.Path != "/" {
		return c.JSON(http.StatusMethodNotAllowed, "not found")
	}
	log.Println(c.Request().Method)
	if c.Request().Method != "GET" {
		return c.JSON(http.StatusMethodNotAllowed, "method not allowed")
	}
	return c.Render(http.StatusOK, "index.html", "")
}
