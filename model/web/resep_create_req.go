package web

type ResepCreateRequest struct {
	Resep       string                     `json:"resep" validate:"required"`
	IdKategori  int                        `json:"id_kategori" validate:"required"`
	ResepDetail []ResepDetailCreateRequest `json:"resep_detail" validate:"required"`
}
