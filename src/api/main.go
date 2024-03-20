package main

import (
	"fmt"
	"log"
	"my-map-api/src/infra/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := Router()
	e.Logger.Fatal(e.Start(":8080"))
}

func Router() *echo.Echo {
	e := echo.New()

	// health check
	e.GET("/", func(c echo.Context) error { return c.JSON(http.StatusOK, "base-api") })

	conn, err := mysql.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn)

	return e
}
