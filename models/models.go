package models

type Todos struct {
	Id          int    //`json:"id" binding:"required, number"`
	Title       string //`json:"tittle" binding:"required"`
	Description string //`json:"description" binding:"required"`
}
