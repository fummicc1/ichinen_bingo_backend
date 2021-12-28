package entity

type Todo struct {
	ID          uint64 `json:"-"`
	Title       string `json:"title"`
	Index       uint64 `json:"index"`
	IsCompleted bool   `json:"is_completed"`
}

type TodoList struct {
	Todos [25]Todo
}
