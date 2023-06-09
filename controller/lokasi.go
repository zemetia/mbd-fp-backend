package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/service"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/umahmood/haversine"
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
	mobilService  service.MobilService
}

func NewLokasiController(us service.LokasiService, ms service.MobilService, jwts service.JWTService) LokasiController {
	return &lokasiController{
		lokasiService: us,
		mobilService:  ms,
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
	longitude, err := strconv.ParseFloat(ctx.DefaultQuery("lon", "106.8456"), 64)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	latitude, err := strconv.ParseFloat(ctx.DefaultQuery("lat", "6.2088"), 64)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	userPos := haversine.Coord{Lat: latitude, Lon: longitude}
	distLimit, err := strconv.ParseFloat(ctx.DefaultQuery("dist", "99999"), 64)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	isPickup, err := strconv.ParseBool(ctx.DefaultQuery("isPickup", "true"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Lokasi", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// Get distances of locations
	var lokasiList []dto.LokasiGetAllDto
	for _, x := range result {
		lat, _ := strconv.ParseFloat(x.Latitude, 64)
		lon, _ := strconv.ParseFloat(x.Longtitude, 64)
		lokasiPos := haversine.Coord{Lat: lat, Lon: lon}
		_, km := haversine.Distance(lokasiPos, userPos)
		if km < distLimit {
			lokasiList = append(lokasiList, dto.LokasiGetAllDto{
				ID:         x.ID,
				Name:       x.Name,
				Longtitude: x.Longtitude,
				Latitude:   x.Latitude,
				Distance:   km,
			})
		}
	}
	// Sort the locations based on distance
	sort.Slice(lokasiList, func(i, j int) bool {
		return lokasiList[i].Distance < lokasiList[j].Distance
	})

	limit := len(lokasiList)
	if limit > 10 {
		limit = 10
	}

	lokasiList = lokasiList[:limit]

	if isPickup {
		res := common.BuildResponse(true, "Berhasil Mendapatkan List Lokasi", lokasiList[:limit])
		ctx.JSON(http.StatusOK, res)
		return
	}

	var listMobil []entity.Mobil
	for _, x := range lokasiList {
		mobil, _ := lc.mobilService.GetMobilByLokasiID(ctx, x.ID)

		if mobil == nil {
			continue
		}

		listMobil = append(listMobil, mobil...)
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan List Mobil", listMobil)
	ctx.JSON(http.StatusOK, res)
}

func (lc *lokasiController) GetLokasi(ctx *gin.Context) {
	lokasiID := ctx.Param("id")

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
	lokasiID := ctx.Param("id")

	err := lc.lokasiService.DeleteLokasi(ctx.Request.Context(), lokasiID)
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

	lokasiID := ctx.Param("id")
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
