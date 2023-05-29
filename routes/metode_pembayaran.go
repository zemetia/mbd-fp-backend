package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

// Todo : Implement Metode Pembayaran Controller
func MetodePembayaranRoutes(router *gin.Engine, MetodePembayaranController controller.MetodePembayaranController, jwtService service.JWTService) {
	metodePembayaranRoutes := router.Group("/api/metode_pembayaran")
	{
		metodePembayaranRoutes.GET("", MetodePembayaranController.GetAllMetodePembayaran)
		metodePembayaranRoutes.POST("", middleware.Authenticate(jwtService), MetodePembayaranController.AddMetodePembayaran)
		metodePembayaranRoutes.DELETE("/:id", middleware.Authenticate(jwtService), MetodePembayaranController.DeleteMetodePembayaran)
		metodePembayaranRoutes.PATCH("/:id", middleware.Authenticate(jwtService), MetodePembayaranController.UpdateMetodePembayaran)
		metodePembayaranRoutes.GET("/:id", middleware.Authenticate(jwtService), MetodePembayaranController.GetMetodePembayaranById)
	}
}
