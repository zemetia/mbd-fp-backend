package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LokasiController interface {
	AddLokasi(ctx *gin.Context)
	GetAllLokasi(ctx *gin.Context)
	GetLokasi(ctx *gin.Context)
	DeleteLokasi(ctx *gin.Context)
	UpdateLokasi(ctx *gin.Context)
}

type lokasiController struct {
	jwtService    service.JWTService
	lokasiService service.LokasiService
}

func NewLokasiController(us service.LokasiService, jwts service.JWTService) LokasiController {
	return &lokasiController{
		lokasiService: us,
		jwtService:    jwts,
	}
}

func (lc *lokasiController) AddLokasi(ctx *gin.Context) {
	var lokasi dto.LokasiCreateDto
	err := ctx.ShouldBind(&lokasi)

	result, err := lc.lokasiService.AddLokasi(ctx.Request.Context(), lokasi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Lokasi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *lokasiController) GetAllLokasi(ctx *gin.Context) {
	result, err := lc.lokasiService.GetAllLokasi(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Lokasi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *lokasiController) GetLokasi(ctx *gin.Context) {
	lokasiID, err := uuid.Parse(ctx.Param("id"))

	result, err := lc.lokasiService.GetLokasi(ctx.Request.Context(), lokasiID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Lokasi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *lokasiController) DeleteLokasi(ctx *gin.Context) {
	lokasiID, err := uuid.Parse(ctx.Param("id"))
	// ctx.Set("token", "")
	// ctx.Set("lokasiID", "")
	err = lc.lokasiService.DeleteLokasi(ctx.Request.Context(), lokasiID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Lokasi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (lc *lokasiController) UpdateLokasi(ctx *gin.Context) {
	var lokasi dto.LokasiUpdateDto
	err := ctx.ShouldBind(&lokasi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	lokasiID, err := uuid.Parse(ctx.Param("id"))
	lokasi.ID = lokasiID
	err = lc.lokasiService.UpdateLokasi(ctx.Request.Context(), lokasi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Lokasi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
