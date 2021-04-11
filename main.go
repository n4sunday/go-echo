package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	var errEnv error
	errEnv = godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error getting env, %v", errEnv)
	}

	e := echo.New()
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(port))
}
