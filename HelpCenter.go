package main

import (
	"./src/controllers"

	"log"
	"net/http"
	"runtime"

	"github.com/facebookgo/grace/gracehttp"
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
	router := gin.New()
	InitMiddleware(router)

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Application, Content-Type",
		ExposedHeaders:  "Content-Length",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := router.Group("/1/notice")
	{
		v1.GET("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
		v1.GET("/list", controllers.ReadListNotice)
		v1.GET("/:id", controllers.ReadIdNotice)
		v1.POST("/add", controllers.CreateNotice)
		v1.PUT("/:id", controllers.UpdateNotice)
		v1.DELETE("/:id", controllers.DeleteNotice)
	}

	v2 := router.Group("/1/faq/category")
	{
		v2.GET("/list", controllers.ReadListFaqCategory)
		v2.GET("/:id", controllers.ReadIdFaqCategory)
		v2.GET("/count/:id", controllers.CountFaqCategory)
		v2.POST("/add", controllers.CreateFaqCategory)
		v2.PUT("/:id", controllers.UpdateFaqCategory)
		v2.DELETE("/:id", controllers.DeleteFaqCategory)
	}

	v3 := router.Group("/1/faq")
	{
		v3.GET("/list", controllers.ReadListFaq)
		v3.GET("/list/:category", controllers.ReadListCategoryFaq)
		v3.GET("/:id", controllers.ReadIdFaq)
		v3.POST("/add", controllers.CreateFaq)
		v3.PUT("/:id", controllers.UpdateFaq)
		v3.DELETE("/:id", controllers.DeleteFaq)
	}

	v4 := router.Group("/1/qna")
	{
		v4.GET("/count", controllers.ReadCountQnA)
		v4.GET("/list", controllers.ReadListQnA)
		v4.GET("/list/:id", controllers.ReadListUserQnA)
		v4.GET("/:id", controllers.ReadIdQnA)
		v4.POST("/add", controllers.CreateQnA)
		v4.POST("/comment/:id", controllers.AddcommentFaq)
		v4.PUT("/:id", controllers.UpdateQnA)
		v4.DELETE("/:id", controllers.DeleteQnA)
	}

	v5 := router.Group("/1/email")
	{
		v5.POST("/send", controllers.SendEmail)
		v5.POST("/export", controllers.ExportEmail)
	}

	err := gracehttp.Serve(
		&http.Server{
			Addr:           ":9090",
			Handler:        router,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		})

	if err != nil {
		log.Println(err)
	}
}
