package main

import (
	"fmt"
	"net/http"
	"trip-planner/db"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal("error loading .env file")
	}

	db, err := db.InitDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println(db)
	// テスト
	e.GET("/api/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
