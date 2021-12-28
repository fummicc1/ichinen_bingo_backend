package entity

type Bingo struct {
	ID       uint64   `json:"id"`
	Title    string   `json:"title"`
	TodoList TodoList `json:"todo_list"`
}
