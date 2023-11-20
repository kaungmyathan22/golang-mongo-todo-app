package models

type ItemModel struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
