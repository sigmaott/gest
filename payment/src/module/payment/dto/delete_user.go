package dto

type DeleteUserById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
