package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/zenklot/catatan-dapur/model/domain"
	"github.com/zenklot/catatan-dapur/model/web"
	"github.com/zenklot/catatan-dapur/repository"
	"gorm.io/gorm"
)

type ResepService interface {
	Create(request web.ResepCreateRequest) web.ResepResponse
	FindById(resepId string) web.ResepDetailResponse
	FindAll() []web.ResepResponse
}

type ResepServiceImpl struct {
	Repository repository.ResepRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewResepService(repository repository.ResepRepository, db *gorm.DB, validate *validator.Validate) *ResepServiceImpl {
	return &ResepServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *ResepServiceImpl) Create(request web.ResepCreateRequest) web.ResepResponse {
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

	id := uuid.New().String()
	resep := domain.Resep{
		Id:         id,
		Resep:      request.Resep,
		IdKategori: request.IdKategori,
	}

	resepDetail := []domain.ResepDetail{}
	for _, bahan := range request.ResepDetail {
		resepDetail = append(resepDetail, domain.ResepDetail{
			IdResep: id,
			IdBahan: bahan.IdBahan,
		})
	}

	resep = service.Repository.Save(tx, resep, resepDetail)

	return web.ResepResponse{}
}

func (service *ResepServiceImpl) FindAll() []web.ResepResponse {
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

	reseps := service.Repository.FindAll(tx)
	var resepResponse []web.ResepResponse
	for _, resep := range reseps {
		resepResponse = append(resepResponse, web.ResepResponse{
			Id:       resep.Id,
			Resep:    resep.Resep,
			Kategori: resep.Kategori.Kategori,
		})
	}

	return resepResponse
}

func (service *ResepServiceImpl) FindById(resepId string) web.ResepDetailResponse {
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

	resep, err := service.Repository.FindById(tx, resepId)
	if err != nil {
		panic(err)
	}

	var bahans []string
	for _, bahan := range resep.ResepDetail {
		bahans = append(bahans, bahan.Bahan.Bahan)
	}

	return web.ResepDetailResponse{
		Id:       resep.Id,
		Resep:    resep.Resep,
		Kategori: resep.Kategori.Kategori,
		Bahan:    bahans,
	}
}
