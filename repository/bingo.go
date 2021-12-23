package repository

import (
	"context"

	"github.com/onyanko-pon/ichinen-bingo/entity"
	"github.com/onyanko-pon/ichinen-bingo/sql_handler"
)

type BingoRepository struct {
	sqlHandler *sql_handler.SQLHandler
}

func NewBingoRepository(sqlHandler *sql_handler.SQLHandler) *BingoRepository {
	return &BingoRepository{
		sqlHandler: sqlHandler,
	}
}

func (r BingoRepository) GetBingo(ctx context.Context, bingoID uint64) (*entity.Bingo, error) {
	query := `SELECT * FROM bingos WHERE bingos.id = $1;`

	rows, err := r.sqlHandler.QueryContext(ctx, query, bingoID)
	if err != nil {
		return nil, err
	}
	var bingo entity.Bingo
	rows.Next()
	err = rows.Scan(&bingo.ID, &bingo.Title)

	if err != nil {
		return nil, err
	}
	return &bingo, nil
}

func (r BingoRepository) Create(ctx context.Context, title string) (*entity.Bingo, error) {
	query := `INSERT INTO bingos (title) VALUES ($1) RETURNING id`

	var bingo entity.Bingo

	rows, err := r.sqlHandler.QueryContext(ctx, query, title)
	rows.Next()
	rows.Scan(&bingo.ID)
	bingo.Title = title

	if err != nil {
		return nil, err
	}

	return &bingo, nil
}

func (r BingoRepository) Update(ctx context.Context, bingo entity.Bingo) (*entity.Bingo, error) {
	query := `UPDATE bingos SET title = $1 WHERE id = $2`

	_, err := r.sqlHandler.QueryContext(ctx, query, bingo.Title, bingo.ID)

	if err != nil {
		return nil, err
	}

	return &bingo, nil
}

func (r BingoRepository) Delete(ctx context.Context, bingoID uint64) error {
	query := `DELETE FROM bingos WHERE id = $1`

	_, err := r.sqlHandler.QueryContext(ctx, query, bingoID)

	return err
}
