package main

import (
	"fp-mbd-amidrive/common"
	"fp-mbd-amidrive/config"
	"fp-mbd-amidrive/controller"
	"fp-mbd-amidrive/middleware"
	"fp-mbd-amidrive/repository"
	"fp-mbd-amidrive/routes"
	"fp-mbd-amidrive/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()

		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)

		tipeRepository repository.TipeRepository = repository.NewTipeRepository(db)
		tipeService    service.TipeService       = service.NewTipeService(tipeRepository)
		tipeController controller.TipeController = controller.NewTipeController(tipeService, jwtService)

		mobilRepository repository.MobilRepository = repository.NewMobilRepository(db)
		mobilService    service.MobilService       = service.NewMobilService(mobilRepository, userRepository, tipeRepository)
		mobilController controller.MobilController = controller.NewMobilController(mobilService, jwtService)

		lokasiRepository repository.LokasiRepository = repository.NewLokasiRepository(db)
		lokasiService    service.LokasiService       = service.NewLokasiService(lokasiRepository)
		lokasiController controller.LokasiController = controller.NewLokasiController(lokasiService, mobilService, jwtService)

		membershipRepository repository.MembershipRepository = repository.NewMembershipRepository(db)
		membershipService    service.MembershipService       = service.NewMembershipService(membershipRepository)
		membershipController controller.MembershipController = controller.NewMembershipController(membershipService, jwtService)

		metodePembayaranRepository repository.MetodePembayaranRepository = repository.NewMetodePembayaranRepository(db)
		metodePembayaranService    service.MetodePembayaranService       = service.NewMetodePembayaranService(metodePembayaranRepository)
		metodePembayaranController controller.MetodePembayaranController = controller.NewMetodePembayaranController(metodePembayaranService, jwtService)

		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)
		transaksiService    service.TransaksiService       = service.NewTransaksiService(transaksiRepository, mobilRepository, userRepository, membershipRepository)
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService, jwtService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.UserRoutes(server, userController, jwtService)
	routes.LokasiRoutes(server, lokasiController, userService, jwtService)
	routes.MobilRoutes(server, mobilController, jwtService)
	routes.MembershipRoutes(server, membershipController, jwtService)
	routes.MetodePembayaranRoutes(server, metodePembayaranController, jwtService)
	routes.TransaksiRoutes(server, transaksiController, jwtService)
	routes.TipeRoutes(server, tipeController, jwtService)

	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	if port == "" {
		port = "8000"
	}
	if ip == "" {
		ip = "localhost"
	}
	server.Run(ip + ":" + port)
}
