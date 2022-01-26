package repository

import (
	"errors"

	"github.com/zenklot/catatan-dapur/model/domain"
	"gorm.io/gorm"
)

type KategoriRepository interface {
	Save(tx *gorm.DB, kategori domain.Kategori) domain.Kategori
	Update(tx *gorm.DB, kategori domain.Kategori) domain.Kategori
	Delete(tx *gorm.DB, kategori domain.Kategori)
	FindById(tx *gorm.DB, kategoriId int) (domain.Kategori, error)
	FindAll(tx *gorm.DB) []domain.Kategori
}

type KategoriRepositoryImpl struct {
}

func NewKategoriRepository() *KategoriRepositoryImpl {
	return &KategoriRepositoryImpl{}
}

func (repository *KategoriRepositoryImpl) Save(tx *gorm.DB, kategori domain.Kategori) domain.Kategori {
	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Create(&kategori)

	if result.Error != nil {
		panic(result.Error)
	}

	return kategori
}

func (repository *KategoriRepositoryImpl) Update(tx *gorm.DB, kategori domain.Kategori) domain.Kategori {
	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Save(&kategori)

	if result.Error != nil {
		panic(result.Error)
	}

	return kategori
}

func (repository *KategoriRepositoryImpl) Delete(tx *gorm.DB, kategori domain.Kategori) {
	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Delete(&kategori)

	if result.Error != nil {
		panic(result.Error)
	}

}

func (repository *KategoriRepositoryImpl) FindById(tx *gorm.DB, kategoriId int) (domain.Kategori, error) {
	if err := tx.Error; err != nil {
		panic(err)
	}

	kategori := domain.Kategori{}
	result := tx.Where("id = ?", kategoriId).First(&kategori)

	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return kategori, errors.New("Id Kategori Tidak Ditemukan")
	} else {
		return kategori, nil
	}
}

func (repository *KategoriRepositoryImpl) FindAll(tx *gorm.DB) []domain.Kategori {
	if err := tx.Error; err != nil {
		panic(err)
	}

	var kategories []domain.Kategori
	result := tx.Find(&kategories)

	if result.Error != nil {
		panic(result.Error)
	}

	return kategories
}
