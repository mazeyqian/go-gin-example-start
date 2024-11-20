package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/mazeyqian/asiatz"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/sites"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func CheckSitesHealth(c *gin.Context) {
	s := persistence.GetRobotRepository()
	webSites, err := getWebSites()
	if err != nil {
		log.Println("error:", err)
		http_err.NewError(c, http.StatusInternalServerError, err)
		return
	}
	markdown, err := s.ClearCheckResult(webSites)
	if err != nil {
		log.Println("error:", err)
		http_err.NewError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": *markdown})
	}
}

func RunCheck() {
	s := persistence.GetRobotRepository()
	// https://github.com/go-co-op/gocron
	// https://pkg.go.dev/time#Location
	UTC, _ := time.LoadLocation("UTC")
	ss := gocron.NewScheduler(UTC)
	everyDayAtStr, _ := asiatz.ShanghaiToUTC("10:00")
	// log.Println("UTC everyDayAtStr:", everyDayAtStr)
	everyDayAtFn := func() {
		sites, err := getWebSites()
		if err != nil {
			log.Println("error:", err)
		} else {
			s.ClearCheckResult(sites)
		}
	}
	ss.Every(1).Day().At(everyDayAtStr).Do(everyDayAtFn)
	ss.StartAsync()
}

func getWebSites() (*[]models.WebSite, error) {
	conf := config.GetConfig()
	webSites := &conf.Data.Sites
	if len(*webSites) == 0 {
		return nil, errors.New("no sites")
	}
	return webSites, nil
}
