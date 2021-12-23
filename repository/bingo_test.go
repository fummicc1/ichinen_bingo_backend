package repository

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/onyanko-pon/ichinen-bingo/sql_handler"
	"github.com/stretchr/testify/assert"
)

var bingoRepository *BingoRepository
var sqlHandler *sql_handler.SQLHandler

func TestMain(m *testing.M) {

	godotenv.Load()
	dataSource := "host=127.0.0.1 port=5432 user=admin password=password dbname=mydb sslmode=disable"
	sqlHandler, _ = sql_handler.NewHandler(dataSource)
	bingoRepository = NewBingoRepository(sqlHandler)
	cleanUp()
	code := m.Run()

	os.Exit(code)
}

func cleanUp() {
	ctx, _ := context.WithCancel(context.Background())
	sqlHandler.CleanData(ctx)
}

func TestCreate(t *testing.T) {
	defer cleanUp()

	title := "title"
	ctx := context.Background()
	bingo, err := bingoRepository.Create(ctx, title)

	assert.Nil(t, err)
	assert.NotNil(t, bingo.ID)
	assert.NotZero(t, bingo.ID)
	assert.Equal(t, title, bingo.Title)
}

func TestGet(t *testing.T) {
	defer cleanUp()
	title := "title"
	ctx := context.Background()
	bingo, _ := bingoRepository.Create(ctx, title)
	id := bingo.ID

	bingo, err := bingoRepository.GetBingo(ctx, bingo.ID)

	assert.Nil(t, err)
	assert.Equal(t, id, bingo.ID)
	assert.Equal(t, title, bingo.Title)
}
