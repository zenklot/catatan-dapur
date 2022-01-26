package web

type ResepUpdateRequest struct {
	Id         int    `json:"id"`
	Resep      string `json:"resep"`
	IdKategori int    `json:"id_kategori"`
}
