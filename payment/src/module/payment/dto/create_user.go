package dto

type CreateUser struct {
	Name string `json:"name" validate:"required"`
}
