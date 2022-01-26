package web

type ResepDetailResponse struct {
	Id       string   `json:"id"`
	Resep    string   `json:"resep"`
	Kategori string   `json:"kategori"`
	Bahan    []string `json:"bahan"`
}
