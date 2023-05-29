package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

func TransaksiRoutes(router *gin.Engine, TransaksiController controller.TransaksiController, jwtService service.JWTService) {
	transaksiRoutes := router.Group("/api/transaksi")
	{
		transaksiRoutes.POST("", middleware.Authenticate(jwtService), TransaksiController.AddTransaksi)
		transaksiRoutes.GET("", middleware.Authenticate(jwtService), TransaksiController.GetAllTransaksi)

		transaksiRoutes.GET("/:id", middleware.Authenticate(jwtService), TransaksiController.GetTransaksi)
		transaksiRoutes.PATCH("/:id", middleware.Authenticate(jwtService), TransaksiController.UpdateTransaksi)
		transaksiRoutes.DELETE("/:id", middleware.Authenticate(jwtService), TransaksiController.DeleteTransaksi)
		//edit mobil status, trigger denda kalo ada
		transaksiRoutes.GET("/:id/kembali", middleware.Authenticate(jwtService), TransaksiController.KembaliMobilTransaksi)
	}
}
