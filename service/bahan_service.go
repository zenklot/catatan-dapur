package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/zenklot/catatan-dapur/model/domain"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/repository"
	"gorm.io/gorm"
)

type BahanService interface {
	Create(request web.BahanCreateRequest) web.BahanResponse
	Update(request web.BahanUpdateRequest) web.BahanResponse
	Delete(bahanId int)
	FindById(bahanId int) web.BahanResponse
	FindAll() []web.BahanResponse
}

type BahanServiceImpl struct {
	BahanRepository repository.BahanRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewBahanService(bahanRepository repository.BahanRepository, db *gorm.DB, vallidate *validator.Validate) *BahanServiceImpl {
	return &BahanServiceImpl{
		BahanRepository: bahanRepository,
		DB:              db,
		Validate:        vallidate,
	}
}

func (service *BahanServiceImpl) Create(request web.BahanCreateRequest) web.BahanResponse {
	if err := service.Validate.Struct(request); err != nil {
		panic(err)
	}

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit().Error; err != nil {
				panic(err)
			}
		}
	}()

	bahan := domain.Bahan{
		Bahan: request.Bahan,
	}

	bahan = service.BahanRepository.Save(tx, bahan)
	return web.BahanResponse{
		Id:    bahan.Id,
		Bahan: bahan.Bahan,
	}
}

func (service *BahanServiceImpl) Update(request web.BahanUpdateRequest) web.BahanResponse {
	panic("not implemented") // TODO: Implement
}

func (service *BahanServiceImpl) Delete(bahanId int) {
	panic("not implemented") // TODO: Implement
}

func (service *BahanServiceImpl) FindById(bahanId int) web.BahanResponse {
	panic("not implemented") // TODO: Implement
}

func (service *BahanServiceImpl) FindAll() []web.BahanResponse {

	tx := service.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			if err := tx.Commit().Error; err != nil {
				panic(err)
			}
		}
	}()

	bahans := service.BahanRepository.FindAll(tx)
	var bahanResponse []web.BahanResponse
	for _, bahan := range bahans {
		bahanResponse = append(bahanResponse, web.BahanResponse{
			Id:    bahan.Id,
			Bahan: bahan.Bahan,
		})
	}
	return bahanResponse
}
