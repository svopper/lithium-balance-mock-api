package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/svopper/lithium-balance-mockapi/structs"
)

// DevicesAll returns a list of devices
func DevicesAll(c *gin.Context) {
	data, _ := ioutil.ReadFile("data/devices-all.json")

	var devices []structs.Device
	err := json.Unmarshal(data, &devices)
	if err != nil {
		c.String(500, "Server error")
		panic(err)
	}

	c.JSON(200, devices)
}
