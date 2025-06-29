package model

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DoneAt      string `json:"done_at"`
	Completed   bool   `json:"completed"`
}
