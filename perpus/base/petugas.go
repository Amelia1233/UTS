package base

import (
	"uts/models"

	"github.com/gin-gonic/gin"
)

type RepositoryP interface {
	GetPetugas(ctx *gin.Context) ([]models.Petugas, error)
	CreatePetugas(data models.Petugas) error
	UpdatePetugas(data models.Petugas, param string) error
	DeletePetugas(param string) error
}

type UsecaseP interface {
	GetPetugas(ctx *gin.Context) ([]models.Petugas, error)
	CreatePetugas(ctx *gin.Context) error
	UpdatePetugas(ctx *gin.Context) error
	DeletePetugas(ctx *gin.Context) error
}
