package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

// Todo : Implement Review Controller
func ReviewRoutes(router *gin.Engine, ReviewController controller.ReviewController, jwtService service.JWTService) {
	reviewRoutes := router.Group("/api/review")
	{
		reviewRoutes.POST("", middleware.Authenticate(jwtService), ReviewController.AddReview)
		reviewRoutes.DELETE("/:id", middleware.Authenticate(jwtService), ReviewController.DeleteReview)
		reviewRoutes.PATCH("/:id", middleware.Authenticate(jwtService), ReviewController.UpdateReview)
		reviewRoutes.GET("/:id", ReviewController.GetReviewById)
		reviewRoutes.GET("/user/:id", ReviewController.GetAllUserReview)
		reviewRoutes.GET("/mobil/:id", ReviewController.GetAllMobilReview)
	}
}
