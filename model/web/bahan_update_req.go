package web

type BahanUpdateRequest struct {
	Id    int    `validate:"required" json:"id"`
	Bahan string `validate:"required,max=200,min=5" json:"bahan"`
}
