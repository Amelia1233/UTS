package usecase

import (
	"log"
	"net/http"
	"uts/models"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func NewUseCasePe(pe base.RepositoryPe) *Usecasest {
	return &Usecasest{
		Repokat: pe,
	}
}

type Usecasest struct {
	Repokat base.RepositoryPe
}

func (us Usecasest) GetPengunjung(ctx *gin.Context) ([]models.Pengunjung, error) {
	result, err := us.Repokat.GetPengunjung(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (us Usecasest) CreatePengunjung(ctx *gin.Context) error {
	var data models.Pengunjung
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = us.Repokat.CreatePengunjung(data)

	if err != nil {
		return err
	}

	return nil
}

func (us Usecasest) UpdatePengunjung(ctx *gin.Context) error {
	var data models.Pengunjung
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}
	ctx.ShouldBindJSON(&data)
	us.Repokat.UpdatePengunjung(data, id)
	return nil
}

func (us Usecasest) DeletePengunjung(ctx *gin.Context) error {
	id := ctx.Param("id")

	if id == "" {
		log.Fatal("ID NOT FOUND")
	}

	err := us.Repokat.DeletePengunjung(id)

	if err != nil {
		return err
	}

	return nil
}
