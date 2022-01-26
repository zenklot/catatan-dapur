package app

import (
	"fmt"

	"github.com/zenklot/catatan-dapur/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {

	dsn := "host=" + Env("DB_HOST") + " user=" + Env("DB_USER") + " password=" + Env("DB_PASSWORD") + " dbname=" + Env("DB_NAME") + " port=" + Env("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&domain.Bahan{}, &domain.ResepDetail{}, &domain.Resep{}, &domain.Kategori{})
	fmt.Println("Database Migrated")
	return DB
}
