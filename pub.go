package main

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"

	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/commonlog"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/cors"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/debug"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/recovery"

	"log"
	"net/http"
	"runtime"

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

	// profiler for gin
	profiler.AddMemoryProfilingHandlers(router)
	// gclogs for gin
	router.GET("/debug/vars", debug.Handler())
	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrapper(router)

	if err := config.Server(router); err != nil {
		log.Println(err)
	}
}
