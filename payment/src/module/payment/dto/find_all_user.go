package dto

type GetListUserQuery struct {
	Q string `json:"q" validate:"required" query:"q"`
}
