package base

import (
	"uts/models"

	"github.com/gin-gonic/gin"
)

type RepositoryBk interface {
	GetBuku(ctx *gin.Context) ([]models.Buku, error)
	CreateBuku(data models.Buku) error
	UpdateBuku(data models.Buku, param string) error
	DeleteBuku(param string) error
}

type UsecaseBk interface {
	GetBuku(ctx *gin.Context) ([]models.Buku, error)
	CreateBuku(ctx *gin.Context) error
	UpdateBuku(ctx *gin.Context) error
	DeleteBuku(ctx *gin.Context) error
}
