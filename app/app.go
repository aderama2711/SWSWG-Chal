package app

import (
	"P1/config"
	"P1/repository"
	"P1/route"
	"P1/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	address := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	router.Run(fmt.Sprintf("%s:%s", address, port))
}
