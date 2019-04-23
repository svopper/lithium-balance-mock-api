package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/svopper/lithium-balance-mockapi/structs"
)

func devicesAll(c *gin.Context) {
	data, _ := ioutil.ReadFile("data/devices-all.json")

	c.String(200, string(data))
}

func deviceStateNow(c *gin.Context) {
	rand := rand.Intn(4)
	var status string

	switch rand {
	case 0:
		status = "Offline"
	case 1:
		status = "Running"
	case 2:
		status = "Idling"
	case 3:
		status = "Error"
	}

	c.String(200, status)
}

func getSite(c *gin.Context) {
	siteID := c.Param("siteId")

	var sites []structs.Site
	sitesJSON, _ := ioutil.ReadFile("data/sites.json")
	err := json.Unmarshal(sitesJSON, &sites)
	if err != nil {
		panic(err)
	}

	var foundSite structs.Site

	for _, site := range sites {
		if site.SiteID == siteID {
			foundSite = site
			break
		}
	}

	c.JSON(200, foundSite)
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.GET("/devices-all", devicesAll)
	router.GET("/devices/:deviceId/states/now", deviceStateNow)
	router.GET("/sites/:siteId", getSite)

	router.Run(":1337")
}
