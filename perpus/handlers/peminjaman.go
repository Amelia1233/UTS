package handlers

import (
	"net/http"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func PeminjamanRouters(pe base.UsecasePm, r *gin.RouterGroup) {
	uc := Usecasestrr{
		pe,
	}

	v2 := r.Group("Peminjaman")
	v2.GET("", uc.GetPeminjaman)
	v2.POST("", uc.CreatePeminjaman)
	v2.PUT(":id", uc.PutPeminjaman)
	v2.DELETE(":id", uc.DeletePeminjaman)
}

type Usecasestrr struct {
	PeUscase base.UsecasePm
}

func (pe Usecasestrr) GetPeminjaman(ctx *gin.Context) {
	result, err := pe.PeUscase.GetPeminjaman(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (pe Usecasestrr) CreatePeminjaman(ctx *gin.Context) {
	err := pe.PeUscase.CreatePeminjaman(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (pe Usecasestrr) PutPeminjaman(ctx *gin.Context) {
	err := pe.PeUscase.UpdatePeminjaman(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (pe Usecasestrr) DeletePeminjaman(ctx *gin.Context) {
	err := pe.PeUscase.DeletePeminjaman(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
