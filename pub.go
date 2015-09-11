package main

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/utility"

	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func handler01(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		next(c)
	})
}

func handler02(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	file, err := os.OpenFile("pub.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	//Create AppContext(rabbitmq)
	// appC := handlers.AppContext{Splk: Splunk.SplunkInit()}
	// defer appC.Close()

	// Gin(Web framework) Release Mode
	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	router := gin.New()

	utility.InitMiddleware(router)
	utility.InitDebuger(router)

	router.GET("/ping", handler01(handler02))

	if err := config.Server(router); err != nil {
		log.Println(err)
	}
}
