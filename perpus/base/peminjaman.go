package base

import (
	"uts/models"

	"github.com/gin-gonic/gin"
)

type RepositoryPm interface {
	GetPeminjaman(ctx *gin.Context) ([]models.Peminjaman, error)
	CreatePeminjaman(data models.Peminjaman) error
	UpdatePeminjaman(data models.Peminjaman, param string) error
	DeletePeminjaman(param string) error
}

type UsecasePm interface {
	GetPeminjaman(ctx *gin.Context) ([]models.Peminjaman, error)
	CreatePeminjaman(ctx *gin.Context) error
	UpdatePeminjaman(ctx *gin.Context) error
	DeletePeminjaman(ctx *gin.Context) error
}
