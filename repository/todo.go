package repository

import (
	"context"

	"github.com/onyanko-pon/ichinen-bingo/entity"
)

func (r BingoRepository) InitTodoList(ctx context.Context, bingoID uint64, todoList entity.TodoList) error {
	query := `INSERT INTO TODOS (title, index, bingo_id, is_completed) values ($1, $2, $3, $4);`

	for _, todo := range todoList.Todos {
		_, err := r.sqlHandler.QueryContext(ctx, query, todo.Title, todo.Index, bingoID, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r BingoRepository) getTodoList(ctx context.Context, bingoID uint64) (*entity.TodoList, error) {
	query := `SELECT id, title, index, is_completed FROM todos WHERE bingo_id = $1 ORDER BY index ASC;`

	rows, err := r.sqlHandler.QueryContext(ctx, query, bingoID)

	if err != nil {
		return nil, err
	}

	todos := [25]entity.Todo{}
	i := 0
	for rows.Next() {
		var todo entity.Todo
		rows.Scan(&todo.ID, &todo.Title, &todo.Index, &todo.IsCompleted)
		todos[i] = todo
		i++
	}

	return &entity.TodoList{Todos: todos}, nil
}

func (r BingoRepository) CompleteTodo(ctx context.Context, bingoID uint64, index uint64) error {
	query := `UPDATE todos is_completed = true WHERE bingo_id = $1 AND index = $2`

	_, err := r.sqlHandler.QueryContext(ctx, query, bingoID, index)
	return err
}
