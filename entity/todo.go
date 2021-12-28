package entity

type Todo struct {
	ID          uint64
	Title       string
	Index       uint64
	IsCompleted bool
}

type TodoList struct {
	Todos [25]Todo
}
