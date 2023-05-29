package routes

import (
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/service"

	"github.com/gin-gonic/gin"
)

// Todo
// Lokasi Controller
// Lokasi: AddLokasi (Name, Longtitude, Latitude)

func LokasiRoutes(router *gin.Engine, LokasiController controller.LokasiController, userService service.UserService, jwtService service.JWTService) {
	lokasiRoutes := router.Group("/api/lokasi")
	{
		lokasiRoutes.POST("", middleware.Authenticate(jwtService), middleware.Role(userService, []string{"admin"}), LokasiController.AddLokasi)
		lokasiRoutes.GET("", middleware.Authenticate(jwtService), LokasiController.GetAllLokasi)

		lokasiRoutes.GET("/:id", middleware.Authenticate(jwtService), LokasiController.GetLokasi)
		lokasiRoutes.PATCH("/:id", middleware.Authenticate(jwtService), middleware.Role(userService, []string{"admin"}), LokasiController.UpdateLokasi)
		lokasiRoutes.DELETE("/:id", middleware.Authenticate(jwtService), middleware.Role(userService, []string{"admin"}), LokasiController.DeleteLokasi)
	}
}
