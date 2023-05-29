package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MobilController interface {
	AddMobil(ctx *gin.Context)
	GetAllMobil(ctx *gin.Context)
	GetMobil(ctx *gin.Context)
	DeleteMobil(ctx *gin.Context)
	UpdateMobil(ctx *gin.Context)
}

type mobilController struct {
	jwtService   service.JWTService
	mobilService service.MobilService
}

func NewMobilController(us service.MobilService, jwts service.JWTService) MobilController {
	return &mobilController{
		mobilService: us,
		jwtService:   jwts,
	}
}

func (lc *mobilController) AddMobil(ctx *gin.Context) {
	var mobil dto.MobilCreateDto
	err := ctx.ShouldBind(&mobil)

	result, err := lc.mobilService.AddMobil(ctx.Request.Context(), mobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *mobilController) GetAllMobil(ctx *gin.Context) {
	result, err := lc.mobilService.GetAllMobil(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *mobilController) GetMobil(ctx *gin.Context) {
	mobilID, err := uuid.Parse(ctx.Param("id"))

	result, err := lc.mobilService.GetMobil(ctx.Request.Context(), mobilID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *mobilController) DeleteMobil(ctx *gin.Context) {
	mobilID, err := uuid.Parse(ctx.Param("id"))
	// ctx.Set("token", "")
	// ctx.Set("mobilID", "")
	err = lc.mobilService.DeleteMobil(ctx.Request.Context(), mobilID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (lc *mobilController) UpdateMobil(ctx *gin.Context) {
	var mobil dto.MobilUpdateDto
	err := ctx.ShouldBind(&mobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	mobilID, err := uuid.Parse(ctx.Param("id"))
	mobil.ID = mobilID
	err = lc.mobilService.UpdateMobil(ctx.Request.Context(), mobil)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
