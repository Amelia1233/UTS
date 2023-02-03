package handlers

import (
	"net/http"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func PengunjungRouters(pe base.UsecasePe, r *gin.RouterGroup) {
	uc := Usecasest{
		pe,
	}

	v2 := r.Group("Pengunjung")
	v2.GET("", uc.GetPengunjung)
	v2.POST("", uc.CreatePengunjung)
	v2.PUT(":id", uc.PutPengunjung)
	v2.DELETE(":id", uc.DeletePengunjung)
}

type Usecasest struct {
	PeUscase base.UsecasePe
}

func (pe Usecasest) GetPengunjung(ctx *gin.Context) {
	result, err := pe.PeUscase.GetPengunjung(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (pe Usecasest) CreatePengunjung(ctx *gin.Context) {
	err := pe.PeUscase.CreatePengunjung(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (pe Usecasest) PutPengunjung(ctx *gin.Context) {
	err := pe.PeUscase.UpdatePengunjung(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (pe Usecasest) DeletePengunjung(ctx *gin.Context) {
	err := pe.PeUscase.DeletePengunjung(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
