package repository

import (
	"errors"

	"github.com/zenklot/catatan-dapur/model/domain"
	"gorm.io/gorm"
)

type ResepRepository interface {
	Save(tx *gorm.DB, resep domain.Resep, detail []domain.ResepDetail) domain.Resep
	FindById(tx *gorm.DB, resepId string) (domain.Resep, error)
	FindAll(tx *gorm.DB) []domain.Resep
}

type ResepRepositoryImpl struct {
}

func NewResepRepository() *ResepRepositoryImpl {
	return &ResepRepositoryImpl{}
}
func (repository *ResepRepositoryImpl) Save(tx *gorm.DB, resep domain.Resep, detail []domain.ResepDetail) domain.Resep {
	if err := tx.Error; err != nil {

		panic(err)
	}

	resepResult := tx.Create(&resep)

	if resepResult.Error != nil {

		panic(resepResult.Error)
	}

	batch := len(detail)
	err := tx.CreateInBatches(detail, batch).Error
	if err != nil {

		panic(err)
	}
	return resep
}

func (repository *ResepRepositoryImpl) FindAll(tx *gorm.DB) []domain.Resep {
	if err := tx.Error; err != nil {
		panic(err)
	}

	var reseps []domain.Resep
	result := tx.Preload("Kategori").Find(&reseps)

	if result.Error != nil {
		panic(result.Error)
	}

	return reseps
}

func (repository *ResepRepositoryImpl) FindById(tx *gorm.DB, resepId string) (domain.Resep, error) {
	if err := tx.Error; err != nil {
		panic(err)
	}

	resep := domain.Resep{
		Id: resepId,
	}
	result := tx.Preload("Kategori").Preload("ResepDetail.Bahan").Find(&resep)

	if result.Error != nil {
		panic(result.Error)
	}

	if result.RowsAffected == 0 {
		return resep, errors.New("Id Resep Tidak Ditemukan")
	} else {
		return resep, nil
	}
}
