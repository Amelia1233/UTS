package repository

import (
	"uts/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repost struct {
	DB *gorm.DB
}

func NewReposPm(Db *gorm.DB) Repost {
	return Repost{
		DB: Db,
	}
}

func (rp Repost) GetPeminjaman(ctx *gin.Context) ([]models.Peminjaman, error) {
	var peminjaman []models.Peminjaman
	rp.DB.Find(&peminjaman)

	return peminjaman, nil
}

func (rp Repost) CreatePeminjaman(data models.Peminjaman) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repost) UpdatePeminjaman(data models.Peminjaman, param string) error {
	err := rp.DB.First(&models.Peminjaman{}, "id=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repost) DeletePeminjaman(param string) error {

	err := rp.DB.First(&models.Peminjaman{}, "id=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Where("id=?",param).Delete(&models.Peminjaman{}).Error

	return nil
}
