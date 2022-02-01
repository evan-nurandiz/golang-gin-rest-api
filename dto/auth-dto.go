package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:6"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
