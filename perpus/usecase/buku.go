package usecase

import (
	"log"
	"net/http"
	"uts/models"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func NewUseCaseBk(pe base.RepositoryBk) *UsecaseBk {
	return &UsecaseBk{
		Repokat: pe,
	}
}

type UsecaseBk struct {
	Repokat base.RepositoryBk
}

func (us UsecaseBk) GetBuku(ctx *gin.Context) ([]models.Buku, error) {
	result, err := us.Repokat.GetBuku(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us UsecaseBk) CreateBuku(ctx *gin.Context) error {
	var data models.Buku
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = us.Repokat.CreateBuku(data)

	if err != nil {
		return err
	}

	return nil
}

func (us UsecaseBk) UpdateBuku(ctx *gin.Context) error {
	var data models.Buku
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data)
	us.Repokat.UpdateBuku(data, id)
	return nil
}

func (us UsecaseBk) DeleteBuku(ctx *gin.Context) error {
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repokat.DeleteBuku(id)

	if err != nil {
		return err
	}

	return nil
}
