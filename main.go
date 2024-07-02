package main

import (
	"MyTransactAPP/config"
	"MyTransactAPP/cron"
	_ "MyTransactAPP/docs"
	"MyTransactAPP/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MyTransactAPP APP
// @version 1.0
// @description This is a sample GO server demonstration of BASIC application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support : Darshan
// @contact.url http://www.swagger.io/support
// @contact.email bdarshan@trellissoft.ai

// @license.name Darshan
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	//load the environment variables
	err := godotenv.Load()
	if err != nil {
		config.Log.Error("Error loading .env file")
	} else {
		config.Log.Info("ENV variables loaded successfully")
	}

	// Initialize the db configuration
	dns := os.Getenv("DATABASE_DSN")
	config.Log.Infof("The database in use details %v", dns)
	DB := config.InitDB(dns)
	r := gin.Default()

	// Setup router
	r = routes.SetupRouter()

	// Schedule the daily report task
	cron.SetupCron(DB)

	r.Run() // listen and serve on 0.0.0.0:8080
}
