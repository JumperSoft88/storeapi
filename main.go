package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func h()  {
	
}

func main() {
	port := os.Getenv("PORT")
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test..")
	})

	//e.Logger.Fatal(e.Start(":1234"))
	e.Logger.Fatal(e.Start(":"+port))
}
