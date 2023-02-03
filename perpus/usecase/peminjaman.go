package usecase

import (
	"log"
	"net/http"
	"uts/models"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func NewUseCasePm(pe base.RepositoryPm) *Usecasetrr {
	return &Usecasetrr{
		Repokat: pe,
	}
}

type Usecasetrr struct {
	Repokat base.RepositoryPm
}

func (us Usecasetrr) GetPeminjaman(ctx *gin.Context) ([]models.Peminjaman, error) {
	result, err := us.Repokat.GetPeminjaman(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecasetrr) CreatePeminjaman(ctx *gin.Context) error {
	var data models.Peminjaman
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = us.Repokat.CreatePeminjaman(data)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecasetrr) UpdatePeminjaman(ctx *gin.Context) error {
	var data models.Peminjaman
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data)
	us.Repokat.UpdatePeminjaman(data, id)
	return nil
}

func (us Usecasetrr) DeletePeminjaman(ctx *gin.Context) error {
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repokat.DeletePeminjaman(id)

	if err != nil {
		return err
	}

	return nil
}
