package usecase

import (
	"log"
	"net/http"
	"uts/models"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func NewUseCaseP(pe base.RepositoryP) *UsecaseP {
	return &UsecaseP{pe}
}

type UsecaseP struct {
	Repokat base.RepositoryP
}

func (us UsecaseP) GetPetugas(ctx *gin.Context) ([]models.Petugas, error) {
	result, err := us.Repokat.GetPetugas(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us UsecaseP) CreatePetugas(ctx *gin.Context) error {
	var data models.Petugas
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = us.Repokat.CreatePetugas(data)

	if err != nil {
		return err
	}

	return nil
}

func (us UsecaseP) UpdatePetugas(ctx *gin.Context) error {
	var data models.Petugas
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data)
	us.Repokat.UpdatePetugas(data, id)
	return nil
}

func (us UsecaseP) DeletePetugas(ctx *gin.Context) error {
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repokat.DeletePetugas(id)

	if err != nil {
		return err
	}

	return nil
}
