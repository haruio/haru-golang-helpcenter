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
)

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

	//Simple group: Upload music
	// v1 := router.Group("/push")
	// {
	// 	v1.GET("/:pushid/ack/:status", appC.Request)
	// }

	//Listen and server on 0.0.0.0:9090
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
