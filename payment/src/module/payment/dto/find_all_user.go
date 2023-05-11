package dto

type GetListUserQuery struct {
	Q string `json:"q" validate:"required,email" query:"q"`
	A string `json:"a"`
}
