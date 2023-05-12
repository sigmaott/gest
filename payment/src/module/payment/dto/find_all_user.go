package dto

type GetListUserQuery struct {
	Q string        `json:"q" validate:"required,email" query:"q"`
	A string        `json:"a" validate:"required,email" query:"q"`
	B ListUserQuery `json:"a" validate:"required"`
}

type ListUserQuery struct {
	Q string `json:"q" validate:"required,email" query:"q"`
	A string `json:"a" validate:"required,email" query:"q"`
}
