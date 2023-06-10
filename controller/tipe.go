package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TipeController interface {
	AddTipeMobil(ctx *gin.Context)
	GetAllTipeMobil(ctx *gin.Context)
	DeleteTipeMobil(ctx *gin.Context)
	UpdateTipeMobil(ctx *gin.Context)

	AddTipeMesin(ctx *gin.Context)
	GetAllTipeMesin(ctx *gin.Context)
	DeleteTipeMesin(ctx *gin.Context)
	UpdateTipeMesin(ctx *gin.Context)

	AddTipePersneling(ctx *gin.Context)
	GetAllTipePersneling(ctx *gin.Context)
	DeleteTipePersneling(ctx *gin.Context)
	UpdateTipePersneling(ctx *gin.Context)
}

type tipeController struct {
	jwtService  service.JWTService
	tipeService service.TipeService
}

func NewTipeController(ts service.TipeService, jwts service.JWTService) TipeController {
	return &tipeController{
		tipeService: ts,
		jwtService:  jwts,
	}
}

func (tc *tipeController) AddTipeMobil(ctx *gin.Context) {
	var tipeMobil dto.CreateTipeMobil
	err := ctx.ShouldBind(&tipeMobil)

	result, err := tc.tipeService.AddTipeMobil(ctx.Request.Context(), tipeMobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Tipe Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) GetAllTipeMobil(ctx *gin.Context) {
	result, err := tc.tipeService.GetAllTipeMobil(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Tipe Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) DeleteTipeMobil(ctx *gin.Context) {
	tipeMobilID, err := strconv.Atoi(ctx.Param("id"))

	err = tc.tipeService.DeleteTipeMobil(ctx.Request.Context(), uint8(tipeMobilID))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Tipe Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) UpdateTipeMobil(ctx *gin.Context) {
	var tipeMobil dto.UpdateTipeMobil
	err := ctx.ShouldBind(&tipeMobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tipeMobilID, err := strconv.Atoi(ctx.Param("id"))
	tipeMobil.ID = uint8(tipeMobilID)
	err = tc.tipeService.UpdateTipeMobil(ctx.Request.Context(), tipeMobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Tipe Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) AddTipeMesin(ctx *gin.Context) {
	var tipeMesin dto.CreateTipeMesin
	err := ctx.ShouldBind(&tipeMesin)

	result, err := tc.tipeService.AddTipeMesin(ctx.Request.Context(), tipeMesin)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Tipe Mesin", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Tipe Mesin", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) GetAllTipeMesin(ctx *gin.Context) {
	result, err := tc.tipeService.GetAllTipeMesin(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Tipe Mesin", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Tipe Mesin", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) DeleteTipeMesin(ctx *gin.Context) {
	tipeMesinID, err := strconv.Atoi(ctx.Param("id"))

	err = tc.tipeService.DeleteTipeMesin(ctx.Request.Context(), uint8(tipeMesinID))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Tipe Mesin", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Tipe Mesin", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) UpdateTipeMesin(ctx *gin.Context) {
	var tipeMesin dto.UpdateTipeMesin
	err := ctx.ShouldBind(&tipeMesin)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Mesin", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tipeMesinID, err := strconv.Atoi(ctx.Param("id"))
	tipeMesin.ID = uint8(tipeMesinID)
	err = tc.tipeService.UpdateTipeMesin(ctx.Request.Context(), tipeMesin)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Mesin", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Tipe Mesin", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) AddTipePersneling(ctx *gin.Context) {
	var tipePersneling dto.CreateTipePersneling
	err := ctx.ShouldBind(&tipePersneling)

	result, err := tc.tipeService.AddTipePersneling(ctx.Request.Context(), tipePersneling)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Tipe Persneling", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Tipe Persneling", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) GetAllTipePersneling(ctx *gin.Context) {
	result, err := tc.tipeService.GetAllTipePersneling(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Tipe Persneling", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Tipe Persneling", result)
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) DeleteTipePersneling(ctx *gin.Context) {
	tipePersnelingID, err := strconv.Atoi(ctx.Param("id"))

	err = tc.tipeService.DeleteTipePersneling(ctx.Request.Context(), uint8(tipePersnelingID))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Tipe Persneling", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Tipe Persneling", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (tc *tipeController) UpdateTipePersneling(ctx *gin.Context) {
	var tipePersneling dto.UpdateTipePersneling
	err := ctx.ShouldBind(&tipePersneling)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Persneling", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	tipePersnelingID, err := strconv.Atoi(ctx.Param("id"))
	tipePersneling.ID = uint8(tipePersnelingID)
	err = tc.tipeService.UpdateTipePersneling(ctx.Request.Context(), tipePersneling)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Tipe Persneling", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Tipe Persneling", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
