package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Hello, world.")
	e := echo.New()
	e.GET("/main", func(c echo.Context) error {
		return c.String(http.StatusOK, "main endpointine GET request yapildi.")
	})
	e.Start(":8080")
}
