package main

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/utility"

	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Create AppContext(rabbitmq)
	// appC := handlers.AppContext{Splk: Splunk.SplunkInit()}
	// defer appC.Close()

	// Gin(Web framework) Release Mode
	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	router := gin.New()

	utility.InitMiddleware(router)
	utility.InitDebug(router)

	if err := config.Server(router); err != nil {
		log.Println(err)
	}
}
