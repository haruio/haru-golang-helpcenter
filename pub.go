package main

import (
	//"./src/database/Splunk"
	"./src/middleware/debug"
	// "./src/handlers"

	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/itsjamie/gin-cors"
	"github.com/DeanThompson/ginpprof"
	"github.com/wblakecaldwell/profiler"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Create AppContext(rabbitmq)
	// appC := handlers.AppContext{Splk: Splunk.SplunkInit()}
	// defer appC.Close()

	// Gin(Web framework) Release Mode
	gin.SetMode(gin.ReleaseMode)

	//Webframework gin-tonic Init
	router := gin.Default()

	//router.Use(gzip.Gzip(gzip.DefaultCompression))	// gzip

	// router.Use(cors.Middleware(cors.Config{		// cors
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// }))

	// v1 := router.Group("/singapore")
	// {
	// 	v1.POST("/screen", appC.Request)
	// }

	router.GET("/ping", func(c *gin.Context) {
		var arr [100]int64
		if arr[0] != 0 {

		}
		c.JSON(http.StatusOK, gin.H{"Status": "ok"})
	})

	// profiler for gin
	profiler.AddMemoryProfilingHandlers(router)
	// gclogs for gin
	router.GET("/debug/vars", debug.Handler())
	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrapper(router)

	//Listen and server on 0.0.0.0:9090
	err := func() error {
		s := &http.Server{
			Addr:           ":9090",
			Handler:        router,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		return s.ListenAndServe()
	}
	if err != nil {
		log.Println(err)
	}
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, router *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, router)
	}

	return http.HandlerFunc(fn)
}
