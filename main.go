package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/svopper/lithium-balance-mockapi/docs"
	"github.com/svopper/lithium-balance-mockapi/structs"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func handleError(c *gin.Context, err error) {
	c.String(500, "Server error")
	panic(err)
}

func devicesAll(c *gin.Context) {
	data, _ := ioutil.ReadFile("data/devices-all.json")

	var devices []structs.Device
	err := json.Unmarshal(data, &devices)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.JSON(200, devices)
}

// /devices/Madeira_test/telemetry/raw?signals=BmsSocTrimmed&last=1
// /devices/Madeira_test/telemetry/raw?signals=InverterActivePower&last=1
func deviceTelementryBmsSocLast(c *gin.Context) {
	signal := c.Query("signals")
	fmt.Println(signal)
	fmt.Println("signal")
	lastNs := c.Query("last")
	var lastN int64
	if len(lastNs) > 0 {
		lastNparsed, strconvErr := strconv.Atoi(lastNs)
		lastN = int64(lastNparsed)
		if strconvErr != nil {
			errorMessage := fmt.Sprintf("Cant parse last arg - %s", lastNs)
			c.String(500, errorMessage)
			return
		}
	} else {
		lastN = 100
	}
	minuteMultiplier := 0
	if signal == "BmsSocTrimmed" {
		var telemetrySlice []structs.BmsSocTrimmed

		for i := int64(0); i < lastN; i++ {
			randI := rand.Intn(100)
			var data structs.BmsSocTrimmed
			data.UTCTime = time.Now().Add(time.Duration(minuteMultiplier) * time.Minute).UTC()
			data.BmsSocTrimmed = int64(randI + 1)

			telemetrySlice = append(telemetrySlice, data)
			minuteMultiplier -= 5
		}

		c.JSON(200, telemetrySlice)
	} else if signal == "InverterActivePower" {
		min := -5.1932296753
		max := 5.1932296753
		var dataSlice []structs.InverterActivePower

		for i := int64(0); i < lastN; i++ {
			randF := min + rand.Float64()*(max-min)
			var data structs.InverterActivePower
			data.UTCTime = time.Now().Add(time.Duration(minuteMultiplier) * time.Minute).UTC()
			data.InverterActivePower = randF

			dataSlice = append(dataSlice, data)
			minuteMultiplier -= 5
		}

		c.JSON(200, dataSlice)
	} else {
		c.String(404, "404, missing signal argument")
	}
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
		c.String(500, "Server error, ")
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

func getSites(c *gin.Context) {
	data, _ := ioutil.ReadFile("data/sites.json")

	var sites []structs.Site
	err := json.Unmarshal(data, &sites)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.JSON(200, sites)
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

// /devices/Madeira_test/telemetry/aggregated?signals=BmsAirInletTemperatureMinCount&signals=BmsAirInletTemperatureMin&signals=BmsAirInletTemperatureMax&signals=BmsAirInletTemperatureMaxCount&last=10
func telemetryAgg(c *gin.Context) {
	// deviceID := c.Param("deviceId")
	signals := c.QueryArray("signals")
	fmt.Println("SIGNALS:")
	fmt.Println(signals)

	if len(signals) == 1 {
		var telemetry []structs.TelemetryAgg
		telemetryJSON, _ := ioutil.ReadFile("data/telemetry-agg.json")
		err := json.Unmarshal(telemetryJSON, &telemetry)
		if err != nil {
			c.String(500, "Server error")
			panic(err)
		}

		c.JSON(200, telemetry)
	} else if len(signals) == 4 {
		var telemetry []structs.TelemetryAggTemp
		telemetryJSON, _ := ioutil.ReadFile("data/temps.json")
		err := json.Unmarshal(telemetryJSON, &telemetry)
		if err != nil {
			c.String(500, "Server error")
			panic(err)
		}

		c.JSON(200, telemetry)
	} else {
		c.String(404, "can't figure out what you mean pl0x s€nd hlep¡")
	}

}

// For swagger setup, see https://github.com/swaggo/swag and https://github.com/swaggo/swag/tree/master/example/celler
// https://lithium-balance-mockapi.herokuapp.com
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://iot-lithiumbalancerm-itu.azurewebsites.net", "http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "OPTION"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authorization"}
	router.Use(cors.New(config))

	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// register routes
	router.GET("/devices-all", devicesAll)
	router.GET("/devices/:deviceId/states/now", deviceStateNow)
	router.GET("/devices/:deviceId/telemetry/raw", deviceTelementryBmsSocLast)
	router.GET("/devices/:deviceId/telemetry/aggregated/signals-all", signalsAll)
	router.GET("/devices/:deviceId/telemetry/aggregated", telemetryAgg)
	router.GET("/sites/:siteId", getSite)
	router.GET("/sites", getSites)

	router.Run()
}
