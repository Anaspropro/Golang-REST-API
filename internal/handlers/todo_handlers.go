package handlers

type CreateTodoInput struct {
	Title string `json:"title"`
	Completed bool `json:"completed"`
}