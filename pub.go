package main

import (
	"./src/database/Splunk"

	"./src/handlers"

	"log"
	//"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	//Create AppContext(rabbitmq)
	appC := handlers.AppContext{Splk: Splunk.SplunkInit()}
	defer appC.Close()

	// Gin(Web framework) Release Mode
	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	r := gin.Default()

	//Simple group: Upload music
	v1 := r.Group("/singapore")
	{
		v1.POST("/screen", appC.Request)
	}

	//Listen and server on 0.0.0.0:9000
	err := r.Run(":9090")
	if err != nil {
		log.Println(err)
	}
}
