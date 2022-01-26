package domain

import "gorm.io/gorm"

type Resep struct {
	gorm.Model
	Id          string        `gorm:"primaryKey;uniqueIndex,not null"`
	Resep       string        `gorm:"not null"`
	IdKategori  int           `gorm:"not null"`
	Kategori    Kategori      `gorm:"foreignKey:IdKategori"`
	ResepDetail []ResepDetail `gorm:"foreignKey:IdResep"`
}
