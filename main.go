package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/svopper/lithium-balance-mockapi/docs"
	"github.com/svopper/lithium-balance-mockapi/structs"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func devicesAll(c *gin.Context) {
	data, _ := ioutil.ReadFile("data/devices-all.json")

	var devices []structs.Device
	err := json.Unmarshal(data, &devices)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.String(200, string(data))
}

// /devices/${id}/telemetry/raw?signals=BmsSocTrimmed&last=1
func deviceTelementryBmsSocLast(c *gin.Context) {
	// signal := c.Query("signals")
	// last := c.Query("last")
	randI := rand.Intn(100)

	BmsJSON, _ := ioutil.ReadFile("data/bms-soc-trimmed.json")

	var telemetryData []structs.BmsSocTrimmed

	err := json.Unmarshal(BmsJSON, &telemetryData)

	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	telemetryData[0].BmsSocTrimmed = int64(randI + 1)
	telemetryData[0].UTCTime = time.Now().UTC()

	c.JSON(200, telemetryData)
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
		c.String(500, "Server error")
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

func signalsAll(c *gin.Context) {
	// deviceID := c.Param("deviceId")
	var signals structs.SignalsAll

	signalsJSON, _ := ioutil.ReadFile("data/signals-all.json")
	err := json.Unmarshal(signalsJSON, &signals)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.JSON(200, signals)
}

func telemetryAgg(c *gin.Context) {
	// deviceID := c.Param("deviceId")
	var telemetry []structs.TelemetryAgg

	telemetryJSON, _ := ioutil.ReadFile("data/telemetry-agg.json")
	err := json.Unmarshal(telemetryJSON, &telemetry)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.JSON(200, telemetry)
}

// For swagger setup, see https://github.com/swaggo/swag and https://github.com/swaggo/swag/tree/master/example/celler
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/devices-all", devicesAll)
	router.GET("/devices/:deviceId/states/now", deviceStateNow)
	router.GET("/devices/:deviceId/telemetry/raw", deviceTelementryBmsSocLast)
	router.GET("/sites/:siteId", getSite)
	router.GET("/devices/:deviceId/telemetry/aggregated/signals-all", signalsAll)
	router.GET("/devices/:deviceId/telemetry/aggregated", telemetryAgg)

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"https://iot-lithiumbalancerm-itu.azurewebsites.net"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.Run()
}
