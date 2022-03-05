package dto

type UserCreateDto struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min=6"`
}

type UserUpdateDTO struct {
	Id       string `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min=6"`
}
