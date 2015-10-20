package main

import (
	"./src/controllers"
	"./src/encoding"
	"./src/models"
	"./src/utility"
	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
	"log"
	// "net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Gin(Web framework) Release Mode
	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
	}))

	//Simple group: Upload music
	v1 := r.Group("/1/notice")
	{
		v1.GET("/:pushid/ack/:status", appC.Request)
		v1.GET("/list", controllers.ReadListNotice)
		v1.GET("/:id", controllers.ReadIdNotice)
		v1.POST("/add", controllers.CreateNotice)
		v1.PUT("/:id",controllers.UpdateNotice)
		v1.DELETE("/:id", controllers.DeleteNotice)
	}

	//Listen and server on 0.0.0.0:9000
	err := r.Run(":10600")
	if err != nil {
		log.Println(err)
	}

	/*
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(encoding.MapEncoder)
	m.Use(models.InitDB())
	m.Use(models.InitRabbitMQ())
	//log.Fatal(http.ListenAndServe(":10600", m))

	f := utility.InitLogger(m)
	defer f.Close()
	*/

//	m.Use(cors.Allow(&cors.Options{
//		AllowOrigins:     []string{"*"},
//		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
//		AllowHeaders:     []string{"Application-Id", "Content-Type"},
//		ExposeHeaders:    []string{"Content-Length"},
//		AllowCredentials: true,
//	}))

	m.Group("/1/notice", func(r martini.Router) {
		r.Get("/list", controllers.ReadListNotice)
		r.Get("/:id", controllers.ReadIdNotice)
		r.Post("/add", binding.Json(controllers.Notice{}), controllers.CreateNotice)
		r.Put("/:id", binding.Json(controllers.Notice{}), controllers.UpdateNotice)
		r.Delete("/:id", binding.Json(controllers.Notice{}), controllers.DeleteNotice)
	})

	m.Group("/1/faq/category", func(r martini.Router) {
		r.Get("/list", controllers.ReadListFaqCategory)
		r.Get("/:id", controllers.ReadIdFaqCategory)
		r.Get("/count/:id", controllers.CountFaqCategory)
		r.Post("/add", binding.Json(controllers.FaqCategory{}), controllers.CreateFaqCategory)
		r.Put("/:id", binding.Json(controllers.FaqCategory{}), controllers.UpdateFaqCategory)
		r.Delete("/:id", binding.Json(controllers.FaqCategory{}), controllers.DeleteFaqCategory)
	})

	m.Group("/1/faq", func(r martini.Router) {
		r.Get("/list", controllers.ReadListFaq)
		r.Get("/list/:category", controllers.ReadListCategoryFaq)
		r.Get("/:id", controllers.ReadIdFaq)
		r.Post("/add", binding.Json(controllers.Faq{}), controllers.CreateFaq)
		r.Put("/:id", binding.Json(controllers.Faq{}), controllers.UpdateFaq)
		r.Delete("/:id", binding.Json(controllers.Faq{}), controllers.DeleteFaq)
	})

	m.Group("/1/qna", func(r martini.Router) {
		r.Get("/count", controllers.ReadCountQnA)
		r.Get("/list", controllers.ReadListQnA)
		r.Get("/list/:id", controllers.ReadListUserQnA)
		r.Get("/:id", controllers.ReadIdQnA)
		r.Post("/add", binding.Json(controllers.QnA{}), controllers.CreateQnA)
		r.Post("/comment/:id", binding.Json(controllers.Comment{}), controllers.AddcommentFaq)
		r.Put("/:id", binding.Json(controllers.QnA{}), controllers.UpdateQnA)
		r.Delete("/:id", binding.Json(controllers.QnA{}), controllers.DeleteQnA)
	})

	m.Group("/1/email", func(r martini.Router) {
		r.Post("/send", binding.Json(controllers.Email{}), controllers.SendEmail)
		r.Post("/export", binding.Json(controllers.MongoExport{}), controllers.ExportEmail)
	})

	m.Run()
}
