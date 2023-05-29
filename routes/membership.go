package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

// Todo : Implement Membership Controller
func MembershipRoutes(router *gin.Engine, MembershipController controller.MembershipController, jwtService service.JWTService) {
	membershipRoutes := router.Group("/api/membership")
	{
		membershipRoutes.GET("", MembershipController.GetAllMembership)
		membershipRoutes.POST("", middleware.Authenticate(jwtService), MembershipController.AddMembership)
		membershipRoutes.GET("/:id", middleware.Authenticate(jwtService), MembershipController.GetMembershipById)
		membershipRoutes.PATCH("/:id", middleware.Authenticate(jwtService), MembershipController.UpdateMembership)
		membershipRoutes.DELETE("/:id", middleware.Authenticate(jwtService), MembershipController.DeleteMembership)
		membershipRoutes.GET("user/:id", middleware.Authenticate(jwtService), MembershipController.GetMembershipByUserId)
	}
}
