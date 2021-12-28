package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/ichinen-bingo/handler"
	"github.com/onyanko-pon/ichinen-bingo/repository"
	"github.com/onyanko-pon/ichinen-bingo/sql_handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var dataSource string
	if os.Getenv("GO_ENV") == "production" {
		dataSource = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)
		fmt.Println(dataSource)
	} else {
		dataSource = "host=127.0.0.1 port=5432 user=admin password=password dbname=mydb sslmode=disable"
	}

	sqlHandler, _ := sql_handler.NewHandler(dataSource)
	bingoRepository := repository.NewBingoRepository(sqlHandler)
	bingoHandler := handler.NewBingoHandler(*bingoRepository)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/bingo/:id", bingoHandler.GetBingo)

	e.Logger.Fatal(
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))),
	)
}
