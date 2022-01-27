package web

type ResepUpdateRequest struct {
	Id       string `json:"id"`
	Resep    string `json:"resep"`
	Kategori int    `json:"kategori"`
	Bahan    []int  `json:"bahan"`
}
