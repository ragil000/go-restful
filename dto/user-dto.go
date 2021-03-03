package dto

// UserUpdateDTO is used by client when PUT /update url
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// UserCreateDTO is used by client when POST /create url
// type UserCreateDTO struct {
// 	Name     string `json:"name" frorm:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
// }
