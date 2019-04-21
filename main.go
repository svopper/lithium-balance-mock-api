package main

import (
	"io/ioutil"
	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func site(c *gin.Context) {
	siteID := c.Param("siteId")

	var data []byte

	switch siteID {
	case "BESS-DK-0000001":
		data, _ = ioutil.ReadFile("data/sites/BESS-DK-0000001.json")
	case "BESS-DK-0000003":
		data, _ = ioutil.ReadFile("data/sites/BESS-DK-0000003.json")
	case "BESS-PT-0000001":
		data, _ = ioutil.ReadFile("data/sites/BESS-PT-0000001.json")
	}

	c.String(200, string(data))
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/devices-all", devicesAll)
	router.GET("/devices/:deviceId/states/now", deviceStateNow)
	router.GET("/sites/:siteId", site)

	router.Run()
}
