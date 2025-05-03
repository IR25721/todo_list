package modles

type Task struct {
	ID  string
	Doc string
}

type List interface {
	TodoList() []Todo
	DoneList() []Done
}
