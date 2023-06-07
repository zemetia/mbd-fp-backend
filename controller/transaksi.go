package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransaksiController interface {
	AddTransaksi(ctx *gin.Context)
	GetAllTransaksi(ctx *gin.Context)
	GetTransaksi(ctx *gin.Context)
	DeleteTransaksi(ctx *gin.Context)
	UpdateTransaksi(ctx *gin.Context)
	KembaliMobilTransaksi(ctx *gin.Context)
}

type transaksiController struct {
	jwtService       service.JWTService
	transaksiService service.TransaksiService
}

func NewTransaksiController(us service.TransaksiService, jwts service.JWTService) TransaksiController {
	return &transaksiController{
		transaksiService: us,
		jwtService:       jwts,
	}
}

func (lc *transaksiController) AddTransaksi(ctx *gin.Context) {
	var transaksi dto.TransaksiCreateDto
	err := ctx.ShouldBind(&transaksi)

	result, err := lc.transaksiService.AddTransaksi(ctx.Request.Context(), transaksi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *transaksiController) GetAllTransaksi(ctx *gin.Context) {
	result, err := lc.transaksiService.GetAllTransaksi(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *transaksiController) GetTransaksi(ctx *gin.Context) {
	transaksiID, err := uuid.Parse(ctx.Param("id"))

	result, err := lc.transaksiService.GetTransaksi(ctx.Request.Context(), transaksiID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Transaksi", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *transaksiController) DeleteTransaksi(ctx *gin.Context) {
	transaksiID, err := uuid.Parse(ctx.Param("id"))
	// ctx.Set("token", "")
	// ctx.Set("transaksiID", "")
	err = lc.transaksiService.DeleteTransaksi(ctx.Request.Context(), transaksiID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Transaksi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (lc *transaksiController) UpdateTransaksi(ctx *gin.Context) {
	var transaksi dto.TransaksiUpdateDto
	err := ctx.ShouldBind(&transaksi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	transaksiID, err := uuid.Parse(ctx.Param("id"))
	transaksi.ID = transaksiID
	err = lc.transaksiService.UpdateTransaksi(ctx.Request.Context(), transaksi)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Transaksi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Transaksi", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (lc *transaksiController) KembaliMobilTransaksi(ctx *gin.Context) {
	transaksiId, err := uuid.Parse(ctx.Param("id"))
	err = lc.transaksiService.KembaliMobilTransaksi(ctx.Request.Context(), transaksiId)

	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengembalikan Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mengembalikan Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
