package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/ichinen-bingo/entity"
	"github.com/onyanko-pon/ichinen-bingo/repository"
)

type BingoHandler struct {
	bingoRepository repository.BingoRepository
}

func NewBingoHandler(bingoRepository repository.BingoRepository) BingoHandler {
	return BingoHandler{
		bingoRepository: bingoRepository,
	}
}

func genBingoEchoMap(bingo entity.Bingo) echo.Map {
	return echo.Map{
		"id":        bingo.ID,
		"title":     bingo.Title,
		"todo_list": bingo.TodoList.Todos,
	}
}

func (handler BingoHandler) GetBingo(ctx echo.Context) error {

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	bingo, _ := handler.bingoRepository.GetBingo(ctx.Request().Context(), uint64(id))

	if bingo == nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": fmt.Sprintf("Bingo Not Found in id %d.", id),
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"bingo": genBingoEchoMap(*bingo),
	})
}

type requestBodyCreateBingo struct {
	Todos []string `json:"todos"`
	Title string   `json:"title"`
}

func (handler BingoHandler) CreateBingo(ctx echo.Context) error {

	requestBody := new(requestBodyCreateBingo)
	if err := ctx.Bind(requestBody); err != nil {
		return err
	}

	if len(requestBody.Todos) != 25 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": fmt.Sprintf("invalid todo length. Expected 24 but %d", len(requestBody.Todos)),
		})
	}

	var todoList entity.TodoList
	for i, todoTitle := range requestBody.Todos {
		todoList.Todos[i] = entity.Todo{
			Title:       todoTitle,
			IsCompleted: false,
			Index:       uint64(i),
		}
	}

	bingo, _ := handler.bingoRepository.Create(ctx.Request().Context(), requestBody.Title, todoList)
	return ctx.JSON(http.StatusOK, echo.Map{
		"bingo": genBingoEchoMap(*bingo),
	})
}
