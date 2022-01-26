package repository

import (
	"errors"

	"github.com/zenklot/catatan-dapur/model/domain"
	"gorm.io/gorm"
)

type BahanRepository interface {
	Save(tx *gorm.DB, bahan domain.Bahan) domain.Bahan
	Update(tx *gorm.DB, bahan domain.Bahan) domain.Bahan
	Delete(tx *gorm.DB, bahan domain.Bahan)
	FindById(tx *gorm.DB, bahanId int) (domain.Bahan, error)
	FindAll(tx *gorm.DB) []domain.Bahan
}

type BahanRepositoryImpl struct {
}

func NewBahanRepository() *BahanRepositoryImpl {
	return &BahanRepositoryImpl{}
}

func (repository *BahanRepositoryImpl) Save(tx *gorm.DB, bahan domain.Bahan) domain.Bahan {

	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Create(&bahan)

	if result.Error != nil {
		panic(result.Error)
	}

	return bahan
}

func (repository *BahanRepositoryImpl) Update(tx *gorm.DB, bahan domain.Bahan) domain.Bahan {
	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Save(&bahan)

	if result.Error != nil {
		panic(result.Error)
	}

	return bahan
}

func (repository *BahanRepositoryImpl) Delete(tx *gorm.DB, bahan domain.Bahan) {
	if err := tx.Error; err != nil {
		panic(err)
	}

	result := tx.Delete(&bahan)

	if result.Error != nil {
		panic(result.Error)
	}

}

func (repository *BahanRepositoryImpl) FindById(tx *gorm.DB, bahanId int) (domain.Bahan, error) {
	if err := tx.Error; err != nil {
		panic(err)
	}

	bahan := domain.Bahan{}
	result := tx.Where("id = ?", bahanId).First(&bahan)

	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected == 0 {
		return bahan, errors.New("Id Bahan Tidak Ditemukan")
	} else {
		return bahan, nil
	}
}

func (repository *BahanRepositoryImpl) FindAll(tx *gorm.DB) []domain.Bahan {
	if err := tx.Error; err != nil {
		panic(err)
	}

	var bahans []domain.Bahan
	result := tx.Find(&bahans)

	if result.Error != nil {
		panic(result.Error)
	}

	return bahans
}
