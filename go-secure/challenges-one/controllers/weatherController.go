package controllers

import (
	"challenges-one/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostWeather(c *gin.Context) {
	water := rand.Intn(100)
	wind := rand.Intn(100)

	// ======================================= WITHOUT STATUS IN JSON =================================== //
	weather_ := models.Weather_{
		Water: water,
		Wind:  wind,
	}

	data_, err_ := json.Marshal(weather_)
	if err_ != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_.Error()})
		return
	}

	resp_, err_ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(string(data_)))
	if err_ != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_.Error()})
		return
	}

	defer resp_.Body.Close()

	body_, err_ := ioutil.ReadAll(resp_.Body)
	if err_ != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err_.Error()})
		return
	}
	fmt.Println(string(body_))
	fmt.Print(getState(water, wind))

	// the result
	// {
	// 	"water": 13,
	// 	"wind": 8,
	//  "id": 101
	// }
	// Status Water : Bahaya
	// Status Wind : Siaga

	// ======================================= WITH STATUS IN JSON =================================== //
	// weather := models.Weather{
	// 	Water:         water,
	// 	Wind:          wind,
	// 	StatusWeather: getState(water, wind),
	// }

	// data, err := json.Marshal(weather)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(string(data)))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Println(string(body))
	// fmt.Println(weather.StatusWeather)

	// the result
	// {
	// 	"water": 13,
	// 	"wind": 8,
	//  "status": "Status Water: bahaya\nStatus Wind: Siaga\n",
	//  "id": 101
	// }
	// Status Water : Bahaya
	// Status Wind : Siaga
}

func getState(water, wind int) string {
	var waterStatus, windStatus string
	if water < 5 {
		waterStatus = "aman"
	} else if water >= 6 && water <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	if wind < 6 {
		windStatus = "aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	return fmt.Sprintf("Status Water : %s\nStatus Wind : %s\n", waterStatus, windStatus)
}
