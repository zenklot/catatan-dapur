package web

type BahanCreateRequest struct {
	Bahan string `validate:"required,max=200,min=4" json:"bahan"`
}
