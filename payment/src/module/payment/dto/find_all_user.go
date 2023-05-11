package dto

type GetListUserQuery struct {
	Q int `json:"q" validate:"required" query:"q"`
	A int `json:"a"`
}
