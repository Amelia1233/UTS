package repository

import (
	"uts/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoBk struct {
	DB *gorm.DB
}

func NewRepoBk(Db *gorm.DB) *RepoBk {
	return &RepoBk{
		DB: Db,
	}
}

func (rp RepoBk) GetBuku(ctx *gin.Context) ([]models.Buku, error) {
	var buku []models.Buku
	rp.DB.Find(&buku)

	return buku, nil
}

func (rp RepoBk) CreateBuku(data models.Buku) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoBk) UpdateBuku(data models.Buku, param string) error {
	err := rp.DB.First(&models.Buku{}, "id=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoBk) DeleteBuku(param string) error {

	err := rp.DB.First(&models.Buku{}, "id=?", param).Error
	if err != nil {
		return err
	}
	err = rp.DB.Where("id=?",param).Delete(&models.Buku{}).Error
	return nil
}
