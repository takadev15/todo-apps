package models

type Todos struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}
