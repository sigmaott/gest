package dto

type GetUserById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
