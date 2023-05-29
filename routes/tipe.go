package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

// Todo : Implement Tipe Controller
func TipeRoutes(router *gin.Engine, TipeController controller.TipeController, jwtService service.JWTService) {
	mobilRoutes := router.Group("/api/tipe/mobil")
	{
		mobilRoutes.GET("", TipeController.GetAllTipeMobil)
		mobilRoutes.POST("", middleware.Authenticate(jwtService), TipeController.AddTipeMobil)
		mobilRoutes.DELETE("/:id", middleware.Authenticate(jwtService), TipeController.DeleteTipeMobil)
		mobilRoutes.PATCH("/:id", middleware.Authenticate(jwtService), TipeController.UpdateTipeMobil)
		mobilRoutes.GET("/:id", TipeController.GetTipeMobilById)
	}

	mesinRoutes := router.Group("/api/tipe/mesin")
	{
		mesinRoutes.GET("", TipeController.GetAllTipeMesin)
		mesinRoutes.POST("", middleware.Authenticate(jwtService), TipeController.AddTipeMesin)
		mesinRoutes.DELETE("/:id", middleware.Authenticate(jwtService), TipeController.DeleteTipeMesin)
		mesinRoutes.PATCH("/:id", middleware.Authenticate(jwtService), TipeController.UpdateTipeMesin)
		mesinRoutes.GET("/:id", TipeController.GetTipeMesinById)
	}

	persnelingRoutes := router.Group("/api/tipe/persneling")
	{
		persnelingRoutes.GET("", TipeController.GetAllTipePersneling)
		persnelingRoutes.POST("", middleware.Authenticate(jwtService), TipeController.AddTipePersneling)
		persnelingRoutes.DELETE("/:id", middleware.Authenticate(jwtService), TipeController.DeleteTipePersneling)
		persnelingRoutes.PATCH("/:id", middleware.Authenticate(jwtService), TipeController.UpdateTipePersneling)
		persnelingRoutes.GET("/:id", TipeController.GetTipePersnelingById)
	}
}
