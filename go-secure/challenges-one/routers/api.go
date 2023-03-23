package routers

import (
	"challenges-one/controllers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	route := gin.Default()

	route.POST("/weathers", controllers.PostWeather)
	go func() {
		for {
			time.Sleep(15 * time.Second)
			_, err := http.Post("http://localhost:8080/weathers", "application/json", nil)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	return route
}
