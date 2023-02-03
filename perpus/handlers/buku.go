package handlers

import (
	"net/http"
	"uts/perpus/base"

	"github.com/gin-gonic/gin"
)

func BukuRouters(pe base.UsecaseBk, r *gin.RouterGroup) {
	uc := Usecases{
		PeUscase: pe,
	}

	v2 := r.Group("Buku")
	v2.GET("", uc.GetBuku)
	v2.POST("", uc.CreateBuku)
	v2.PUT(":id", uc.PutBuku)
	v2.DELETE(":id", uc.DeleteBuku)
}

type Usecases struct {
	PeUscase base.UsecaseBk
}

func (pe Usecases) GetBuku(ctx *gin.Context) {
	result, err := pe.PeUscase.GetBuku(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (pe Usecases) CreateBuku(ctx *gin.Context) {
	err := pe.PeUscase.CreateBuku(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (pe Usecases) PutBuku(ctx *gin.Context) {
	err := pe.PeUscase.UpdateBuku(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (pe Usecases) DeleteBuku(ctx *gin.Context) {
	err := pe.PeUscase.DeleteBuku(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
