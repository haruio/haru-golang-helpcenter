package main

import (
	"./src/controllers"
	"./src/encoding"
	"./src/models"
	"./src/utility"

	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/commonlog"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/recovery"
)

func InitMiddleware(router *gin.Engine) {
	router.Use(commonlog.Logger())  // logger
	router.Use(recovery.Recovery()) // recover
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Create AppContext(rabbitmq)
	appC := handlers.AppContext{RabbitMQ.PublisherInit()}
	defer RabbitMQ.Close()

	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Application, Content-Type",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	s := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
