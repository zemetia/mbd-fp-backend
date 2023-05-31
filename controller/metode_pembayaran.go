package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MetodePembayaranController interface {
	AddMetodePembayaran(ctx *gin.Context)
	GetAllMetodePembayaran(ctx *gin.Context)
	GetMetodePembayaran(ctx *gin.Context)
	DeleteMetodePembayaran(ctx *gin.Context)
	UpdateMetodePembayaran(ctx *gin.Context)
}

type metodePembayaranController struct {
	jwtService              service.JWTService
	metodePembayaranService service.MetodePembayaranService
}

func NewMetodePembayaranController(us service.MetodePembayaranService, jwts service.JWTService) MetodePembayaranController {
	return &metodePembayaranController{
		metodePembayaranService: us,
		jwtService:              jwts,
	}
}

func (lc *metodePembayaranController) AddMetodePembayaran(ctx *gin.Context) {
	var metodePembayaran dto.MetodePembayaranCreateDto
	err := ctx.ShouldBind(&metodePembayaran)

	result, err := lc.metodePembayaranService.AddMetodePembayaran(ctx.Request.Context(), metodePembayaran)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan MetodePembayaran", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *metodePembayaranController) GetAllMetodePembayaran(ctx *gin.Context) {
	result, err := lc.metodePembayaranService.GetAllMetodePembayaran(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List MetodePembayaran", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *metodePembayaranController) GetMetodePembayaran(ctx *gin.Context) {
	metodePembayaranID, err := uuid.Parse(ctx.Param("id"))

	result, err := lc.metodePembayaranService.GetMetodePembayaran(ctx.Request.Context(), metodePembayaranID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List MetodePembayaran", result)
	ctx.JSON(http.StatusOK, res)
}

func (lc *metodePembayaranController) DeleteMetodePembayaran(ctx *gin.Context) {
	metodePembayaranID, err := uuid.Parse(ctx.Param("id"))
	// ctx.Set("token", "")
	// ctx.Set("metodePembayaranID", "")
	err = lc.metodePembayaranService.DeleteMetodePembayaran(ctx.Request.Context(), metodePembayaranID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus MetodePembayaran", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (lc *metodePembayaranController) UpdateMetodePembayaran(ctx *gin.Context) {
	var metodePembayaran dto.MetodePembayaranUpdateDto
	err := ctx.ShouldBind(&metodePembayaran)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	metodePembayaranID, err := uuid.Parse(ctx.Param("id"))
	metodePembayaran.ID = metodePembayaranID
	err = lc.metodePembayaranService.UpdateMetodePembayaran(ctx.Request.Context(), metodePembayaran)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate MetodePembayaran", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate MetodePembayaran", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
