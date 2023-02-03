package router

import (
	"uts/perpus/handlers"
	"uts/perpus/repository"
	"uts/perpus/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerEndPoint struct {
	DB *gorm.DB
	R  *gin.Engine
}

func (HEP HandlerEndPoint) Routers() {
	v1 := HEP.R.Group("Api")
	//Petugas
	RepoP := repository.NewRepoP(HEP.DB)
	UsecaseP := usecase.NewUseCaseP(RepoP)
	handlers.PetugasRouters(UsecaseP, v1)

	//Buku
	RepoBk := repository.NewRepoBk(HEP.DB)
	UsecaseBk := usecase.NewUseCaseBk(RepoBk)
	handlers.BukuRouters(UsecaseBk, v1)

	//Pengunjung
	RepoPe := repository.NewRepoPe(HEP.DB)
	UsecasePe := usecase.NewUseCasePe(RepoPe)
	handlers.PengunjungRouters(UsecasePe, v1)

	//Peminjaman
	RepoPem := repository.NewReposPm(HEP.DB)
	UsePem := usecase.NewUseCasePm(RepoPem)
	handlers.PeminjamanRouters(UsePem, v1)
}
