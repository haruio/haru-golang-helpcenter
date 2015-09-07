package utility

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/commonlog"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/cors"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/debug"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/recovery"

	"log"
	"net/http"

	"github.com/DeanThompson/ginpprof"
	"github.com/HyeJong/profiler"
	"github.com/gin-gonic/gin"
)

func ParamCheck(key string, c *gin.Context) string {
	param := c.Param(key)
	if param == "" {
		c.JSON(http.StatusPaymentRequired, gin.H{"Statues": "Require parameters"})
		log.Panic("Require parameters")
	}

	return param
}

func GinErrCheck(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": err})
		log.Panic(err.Error())
	}
}

func ErrCheck(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func InitMiddleware(router *gin.Engine) {
	router.Use(commonlog.Logger())
	router.Use(recovery.Recovery())
	router.Use(cors.Middleware(config.CORS_CONFIG))
	//router.Use(gzip.Gzip(gzip.DefaultCompression))	// gzip
}

func InitDebug(router *gin.Engine) {

	// gclogs for gin
	router.GET("/debug/vars", debug.Handler())

	// profiler for gin
	profiler.AddMemoryProfilingHandlers(router)

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrapper(router)
}
