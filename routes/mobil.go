package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

func MobilRoutes(router *gin.Engine, MobilController controller.MobilController, jwtService service.JWTService) {
	mobilRoutes := router.Group("/api/mobil")
	{
		mobilRoutes.POST("", middleware.Authenticate(jwtService), MobilController.AddMobil)
		mobilRoutes.GET("", middleware.Authenticate(jwtService), MobilController.GetAllMobil)

		mobilRoutes.GET("/:id", middleware.Authenticate(jwtService), MobilController.GetMobil)
		mobilRoutes.PATCH("/:id", middleware.Authenticate(jwtService), MobilController.UpdateMobil)
		mobilRoutes.DELETE("/:id", middleware.Authenticate(jwtService), MobilController.DeleteMobil)
	}
}
