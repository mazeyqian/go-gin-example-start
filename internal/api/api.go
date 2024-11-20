package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/api/controllers"
	"github.com/mazeyqian/go-gin-gee/internal/api/router"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/db"
)

func setConfiguration() {
	config.Setup()
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run() {
	// Set the timezone to UTC
	// https://www.zeitverschiebung.net/en/timezone/asia--shanghai
	os.Setenv("TZ", "UTC")
	setConfiguration()
	conf := config.GetConfig()
	log.Println("Config:", conf)
	// Run before the API starts
	if len(conf.Data.Sites) > 0 {
		// log.Println("check for sites:", conf.Data.Sites)
		controllers.RunCheck()
	} else {
		log.Println("no sites, unnecessary to run check")
	}
	web := router.Setup()
	fmt.Println("Go API Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
