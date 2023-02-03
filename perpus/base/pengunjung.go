package base

import (
	"uts/models"

	"github.com/gin-gonic/gin"
)

type RepositoryPe interface {
	GetPengunjung(ctx *gin.Context) ([]models.Pengunjung, error)
	CreatePengunjung(data models.Pengunjung) error
	UpdatePengunjung(data models.Pengunjung, param string) error
	DeletePengunjung(param string) error
}

type UsecasePe interface {
	GetPengunjung(ctx *gin.Context) ([]models.Pengunjung, error)
	CreatePengunjung(ctx *gin.Context) error
	UpdatePengunjung(ctx *gin.Context) error
	DeletePengunjung(ctx *gin.Context) error
}
