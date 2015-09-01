package main

import (
	//"./src/database/Splunk"
	"./src/handlers/activity_handler"

	"./src"

	"./src/middleware/commonlog"
	"./src/middleware/cors"
	"./src/middleware/debug"
	"./src/middleware/recovery"

	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/HyeJong/profiler"
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

	router.Use(commonlog.Logger())
	router.Use(recovery.Recovery())
	router.Use(cors.Middleware(config.CORS_CONFIG))
	//router.Use(gzip.Gzip(gzip.DefaultCompression))	// gzip

	v1 := router.Group("/singapore")
	{
		v1.POST("/screen", activity_handler.Request)
	}

	// profiler for gin
	profiler.AddMemoryProfilingHandlers(router)
	// gclogs for gin
	router.GET("/debug/vars", debug.Handler())
	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrapper(router)

	// Listen and server on 0.0.0.0:9090
	ser := func() error {
		s := &http.Server{
			Addr:           ":9090",
			Handler:        router,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		return s.ListenAndServe()
	}

	if err := ser(); err != nil {
		log.Println(err)
	}
}
