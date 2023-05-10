package dto

type UpdateUser struct {
	ID   string `json:"id" validate:"required" param:"id"`
	Name string `json:"name" validate:"required"`
}
