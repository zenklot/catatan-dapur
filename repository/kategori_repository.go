package repository

import (
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

func NewKategoryRepository() *KategoriRepositoryImpl {
	return &KategoriRepositoryImpl{}
}

func (repository *KategoriRepositoryImpl) Save(tx *gorm.DB, kategori domain.Kategori) domain.Kategori {
	panic("not implemented") // TODO: Implement
}

func (repository *KategoriRepositoryImpl) Update(tx *gorm.DB, kategori domain.Kategori) domain.Kategori {
	panic("not implemented") // TODO: Implement
}

func (repository *KategoriRepositoryImpl) Delete(tx *gorm.DB, kategori domain.Kategori) {
	panic("not implemented") // TODO: Implement
}

func (repository *KategoriRepositoryImpl) FindById(tx *gorm.DB, kategoriId int) (domain.Kategori, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *KategoriRepositoryImpl) FindAll(tx *gorm.DB) []domain.Kategori {
	panic("not implemented") // TODO: Implement
}
