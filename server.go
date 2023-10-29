package main

import (
	"fmt"
	"net/http"
	// "log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func main(){
	e := echo.New()

	DBInit()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	fmt.Println("XXXXXXXXX", os.Getenv("CLUSTER_USERNAME"))

	e.Logger.Fatal(e.Start(":9001"))
}

func hello(c echo.Context) error {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	return c.JSONPretty(http.StatusOK, data, "  ")

}

