package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewController interface {
	AddReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	GetReviewById(ctx *gin.Context)
	GetAllMobilReview(ctx *gin.Context)
	GetAllUserReview(ctx *gin.Context)
}

type reviewController struct {
	jwtService    service.JWTService
	reviewService service.ReviewService
}

func NewReviewController(rs service.ReviewService, jwts service.JWTService) ReviewController {
	return &reviewController{
		reviewService: rs,
		jwtService:    jwts,
	}
}

func (rc *reviewController) AddReview(ctx *gin.Context) {
	var reviewDTO dto.ReviewCreateDto
	err := ctx.ShouldBind(&reviewDTO)

	result, err := rc.reviewService.AddReview(ctx.Request.Context(), reviewDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Review Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Review Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (rc *reviewController) DeleteReview(ctx *gin.Context) {
	reviewID, err := uuid.Parse(ctx.Param("id"))

	err = rc.reviewService.DeleteReview(ctx.Request.Context(), reviewID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Review Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Review Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (rc *reviewController) UpdateReview(ctx *gin.Context) {
	var reviewDTO dto.ReviewUpdateDto
	err := ctx.ShouldBind(&reviewDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Review Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reviewID, err := uuid.Parse(ctx.Param("id"))
	reviewDTO.ID = reviewID
	err = rc.reviewService.UpdateReview(ctx.Request.Context(), reviewDTO)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Review Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Review Mobil", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (rc *reviewController) GetReviewById(ctx *gin.Context) {
	reviewID, err := uuid.Parse(ctx.Param("id"))

	result, err := rc.reviewService.GetReviewById(ctx.Request.Context(), reviewID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Review Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Review Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (rc *reviewController) GetAllMobilReview(ctx *gin.Context) {
	mobilID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := rc.reviewService.GetAllMobilReview(ctx.Request.Context(), mobilID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Tipe Mobil", result)
	ctx.JSON(http.StatusOK, res)
}

func (rc *reviewController) GetAllUserReview(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := rc.reviewService.GetAllMobilReview(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Tipe Mobil", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Tipe Mobil", result)
	ctx.JSON(http.StatusOK, res)
}
