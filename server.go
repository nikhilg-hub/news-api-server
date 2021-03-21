package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/topHeadlines", handler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func handler(c echo.Context) error {

	country := c.QueryParam("country")
	apiKey := c.QueryParam("apiKey")

	resp, err := http.Get("http://newsapi.org/v2/top-headlines?country=" +
		country + "&apiKey=" + apiKey)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Error")
	} else {
		body, _ := io.ReadAll(resp.Body)
		return c.String(http.StatusOK, string(body))
	}
}
