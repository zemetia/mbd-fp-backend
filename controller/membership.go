package controller

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MembershipController interface {
	GetAllMembership(ctx *gin.Context)
	AddMembership(ctx *gin.Context)
	UpdateMembership(ctx *gin.Context)
	DeleteMembership(ctx *gin.Context)
	GetMembershipByUserId(ctx *gin.Context)
	GetMembershipById(ctx *gin.Context)
}

type membershipController struct {
	jwtService        service.JWTService
	membershipService service.MembershipService
}

func NewMembershipController(ms service.MembershipService, jwts service.JWTService) MembershipController {
	return &membershipController{
		membershipService: ms,
		jwtService:        jwts,
	}
}

func (ms *membershipController) GetAllMembership(ctx *gin.Context) {
	result, err := ms.membershipService.GetAllMembership(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan List Membership", result)
	ctx.JSON(http.StatusOK, res)
}

func (ms *membershipController) AddMembership(ctx *gin.Context) {
	var membership dto.MembershipCreateDto
	err := ctx.ShouldBind(&membership)

	result, err := ms.membershipService.AddMembership(ctx.Request.Context(), membership)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Membership", result)
	ctx.JSON(http.StatusOK, res)
}

func (ms *membershipController) UpdateMembership(ctx *gin.Context) {
	var membership dto.MembershipUpdateDto
	err := ctx.ShouldBind(&membership)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menambahkan Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	membershipTier := ctx.Param("id")
	membership.Tier = membershipTier

	err = ms.membershipService.UpdateMembership(ctx.Request.Context(), membership)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mengupdate Membership", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (ms *membershipController) DeleteMembership(ctx *gin.Context) {
	err := ms.membershipService.DeleteMembership(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menghapus Membership", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (ms *membershipController) GetMembershipByUserId(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Membership User", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := ms.membershipService.GetMembershipByUserId(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Membership User", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Membership User", result)
	ctx.JSON(http.StatusOK, res)
}

func (ms *membershipController) GetMembershipById(ctx *gin.Context) {
	result, err := ms.membershipService.GetMembershipById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Membership", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mendapatkan Membership", result)
	ctx.JSON(http.StatusOK, res)
}
