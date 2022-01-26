package domain

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Id       int     `gorm:"primaryKey;not null"`
	Kategori string  `gorm:"not null"`
	Resep    []Resep `gorm:"foreignKey:IdKategori"`
}
