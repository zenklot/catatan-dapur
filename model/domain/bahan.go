package domain

import "gorm.io/gorm"

type Bahan struct {
	gorm.Model
	Id          int           `gorm:"primaryKey;not null;autoIncrement"`
	Bahan       string        `gorm:"not null"`
	ResepDetail []ResepDetail `gorm:"foreignKey:IdBahan"`
}
