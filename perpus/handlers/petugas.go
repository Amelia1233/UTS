package handlers

import (
	"net/http"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func PetugasRouters(pe base.UsecaseP, r *gin.RouterGroup) {
	uc := Usecasestr{
		pe,
	}

	v2 := r.Group("Petugas")
	v2.GET("", uc.GetPetugas)
	v2.POST("", uc.CreatePetugas)
	v2.PUT(":id", uc.PutPetugas)
	v2.DELETE(":id", uc.DeletePetugas)
}

type Usecasestr struct {
	PeUscase base.UsecaseP
}

func (pe Usecasestr) GetPetugas(ctx *gin.Context) {
	result, err := pe.PeUscase.GetPetugas(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (pe Usecasestr) CreatePetugas(ctx *gin.Context) {
	err := pe.PeUscase.CreatePetugas(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (pe Usecasestr) PutPetugas(ctx *gin.Context) {
	err := pe.PeUscase.UpdatePetugas(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (pe Usecasestr) DeletePetugas(ctx *gin.Context) {
	err := pe.PeUscase.DeletePetugas(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
