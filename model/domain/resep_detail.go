package domain

import "gorm.io/gorm"

type ResepDetail struct {
	gorm.Model
	Id      int    `gorm:"primaryKey;not null"`
	IdResep string `gorm:"not null"`
	Resep   Resep  `gorm:"foreignKey:IdResep"`
	IdBahan int    `gorm:"not null"`
	Bahan   Bahan  `gorm:"foreignKey:IdBahan"`
}
