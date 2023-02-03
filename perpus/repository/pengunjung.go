package repository

import (
	"uts/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Reposs struct {
	DB *gorm.DB
}

func NewRepoPe(Db *gorm.DB) Reposs {
	return Reposs{
		DB: Db,
	}
}

func (rp Reposs) GetPengunjung(ctx *gin.Context) ([]models.Pengunjung, error) {
	var pengunjung []models.Pengunjung
	rp.DB.Find(&pengunjung)

	return pengunjung, nil
}

func (rp Reposs) CreatePengunjung(data models.Pengunjung) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Reposs) UpdatePengunjung(data models.Pengunjung, param string) error {
	err := rp.DB.First(&models.Pengunjung{}, "id=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Reposs) DeletePengunjung(param string) error {

	err := rp.DB.First(&models.Pengunjung{}, "id=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Where("id=?",param).Delete(&models.Pengunjung{}).Error

	return nil
}
