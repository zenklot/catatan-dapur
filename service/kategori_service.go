package service

import (
	"github.com/go-playground/validator/v10"
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

func NewKategoriServiceImpl(repository repository.KategoriRepository, db *gorm.DB, validate *validator.Validate) *KategoriServiceImpl {
	return &KategoriServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *KategoriServiceImpl) Create(request web.KategoriCreateRequest) web.KategoriResponse {
	panic("not implemented") // TODO: Implement
}

func (service *KategoriServiceImpl) Update(request web.KategoriUpdateRequest) web.KategoriResponse {
	panic("not implemented") // TODO: Implement
}

func (service *KategoriServiceImpl) Delete(kategoriId int) {
	panic("not implemented") // TODO: Implement
}

func (service *KategoriServiceImpl) FindById(kategoriId int) web.KategoriResponse {
	panic("not implemented") // TODO: Implement
}

func (service *KategoriServiceImpl) FindAll() []web.KategoriResponse {
	panic("not implemented") // TODO: Implement
}
