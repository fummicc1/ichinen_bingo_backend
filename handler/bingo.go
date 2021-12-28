package handler

import (
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

func (handler BingoHandler) GetBingo(ctx echo.Context) error {

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	bingo, _ := handler.bingoRepository.GetBingo(ctx.Request().Context(), uint64(id))
	return ctx.JSON(http.StatusOK, echo.Map{
		"bingo": bingo,
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
			"message": "invalid todo length",
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
		"bingo": bingo,
	})
}
