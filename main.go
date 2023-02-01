package main

import (
	"Haioo_6/api"
	"Haioo_6/api/controller"
	"Haioo_6/api/injection"
	"Haioo_6/api/service"
	"Haioo_6/api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

var (
	db             *gorm.DB                  = injection.CreateDatabase()
	Migration      injection.Migration       = injection.NewMigration(db)
	CartService    service.CartService       = service.NewCartService(db)
	CartController controller.CartController = controller.NewCartController(CartService)
	Routes         api.Route                 = api.NewRoute(CartController)
)

func main() {
	defer injection.CloseDatabaseConnection(db)

	Migration.Migrate()

	router := gin.Default()

	Routes.Routes(router)

	port := utils.EnvVar("PORT", "8080")
	err := router.Run(":" + port)
	if err != nil {
		log.Println("Failed To Start System")
	}
}
