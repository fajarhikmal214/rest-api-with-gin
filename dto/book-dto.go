package dto

type BookCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      string `json:"user_id" form:"user_id" binding:"required"`
}

type BookUpdateDTO struct {
	Id          string `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId      string `json:"user_id" form:"user_id" binding:"required"`
}
