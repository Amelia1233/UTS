package repository

import (
	"uts/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepoP(Db *gorm.DB) *Repo {
	return &Repo{
		DB: Db,
	}
}

func (rp Repo) GetPetugas(ctx *gin.Context) ([]models.Petugas, error) {
	var petugas []models.Petugas
	rp.DB.Find(&petugas)

	return petugas, nil
}

func (rp Repo) CreatePetugas(data models.Petugas) error {
	err := rp.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo) UpdatePetugas(data models.Petugas, param string) error {
	err := rp.DB.First(&models.Petugas{}, "id=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp Repo) DeletePetugas(param string) error {

	err := rp.DB.First(&models.Petugas{}, "id=?", param).Error
	if err != nil {
		return err
	}

	err = rp.DB.Where("id=?",param).Delete(&models.Petugas{}).Error

	return nil
}
