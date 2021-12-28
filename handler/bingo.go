package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
