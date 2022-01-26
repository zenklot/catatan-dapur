package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/zenklot/catatan-dapur/app/exception"
	"github.com/zenklot/catatan-dapur/model/domain"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/repository"
	"gorm.io/gorm"
)

type KategoriService interface {
	Create(request web.KategoriCreateRequest) web.KategoriResponse
	Update(request web.KategoriUpdateRequest) web.KategoriResponse
	Delete(kategoriId int)
	FindById(kategoriId int) web.KategoriResponse
	FindAll() []web.KategoriResponse
}

type KategoriServiceImpl struct {
	Repository repository.KategoriRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewKategoriService(repository repository.KategoriRepository, db *gorm.DB, validate *validator.Validate) *KategoriServiceImpl {
	return &KategoriServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *KategoriServiceImpl) Create(request web.KategoriCreateRequest) web.KategoriResponse {
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

	kategori := domain.Kategori{
		Kategori: request.Kategori,
	}

	kategori = service.Repository.Save(tx, kategori)
	return web.KategoriResponse{
		Id:       kategori.Id,
		Kategori: kategori.Kategori,
	}

}

func (service *KategoriServiceImpl) Update(request web.KategoriUpdateRequest) web.KategoriResponse {
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

	kategori := domain.Kategori{
		Id:       request.Id,
		Kategori: request.Kategori,
	}

	kategori = service.Repository.Update(tx, kategori)
	return web.KategoriResponse{
		Id:       kategori.Id,
		Kategori: kategori.Kategori,
	}
}

func (service *KategoriServiceImpl) Delete(kategoriId int) {

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

	kategori := domain.Kategori{
		Id: kategoriId,
	}

	service.Repository.Delete(tx, kategori)

}

func (service *KategoriServiceImpl) FindById(kategoriId int) web.KategoriResponse {

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

	kategori, err := service.Repository.FindById(tx, kategoriId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.KategoriResponse{
		Id:       kategori.Id,
		Kategori: kategori.Kategori,
	}
}

func (service *KategoriServiceImpl) FindAll() []web.KategoriResponse {

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

	kategories := service.Repository.FindAll(tx)
	var kategoriResponse []web.KategoriResponse
	for _, kategori := range kategories {
		kategoriResponse = append(kategoriResponse, web.KategoriResponse{
			Id:       kategori.Id,
			Kategori: kategori.Kategori,
		})
	}

	return kategoriResponse
}
