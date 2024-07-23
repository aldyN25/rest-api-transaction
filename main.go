package main

import (
	"fmt"
	"log"

	"github.com/aldyN25/myproject/app/config"
	apiV1 "github.com/aldyN25/myproject/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	apiV1.SetupRouter(router)
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	config := config.GetInstance()

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	log.Println("LISTEN =============== ", listen)
	router.Run(listen)
}
