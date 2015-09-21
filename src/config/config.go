// config
package config

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/cors"

	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

const (
	NAMESPACE = "st"

	SPLUNK_PROTOCOL = "tcp"
	SPLUNK_ADDR     = "127.0.0.1:9998"

	MONGODB_ADDR = "127.0.0.1:27017"

	RDB_TYPE = "mysql"
	RDB_ADDR = ""
	// The Data Source Name has a common format, like e.g. PEAR DB uses it, but without type-prefix
	// 	- [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// example
	// 	- user@unix(/path/to/socket)/dbname?charset=utf8
	// 	- user:password@tcp(localhost:5555)/dbname?charset=utf8
	// 	- user:password@/dbname
	// 	- user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	COUCHBASE_ADDR       = "couchbase://127.0.0.1"
	COUCHBASE_BUCKETNAME = "beer-sample"
	COUCHBASE_PASSWORD   = ""

	RABBITMQ_ADDR      = "amqp://guest:guest@localhost:5672/"
	RABBITMQ_QUEUENAME = "Queue"
)

var CORS_CONFIG = cors.Config{
	Origins:         "*",
	Methods:         "GET, PUT, POST, DELETE",
	RequestHeaders:  "Origin, Authorization, Content-Type",
	ExposedHeaders:  "",
	MaxAge:          50 * time.Second,
	Credentials:     true,
	ValidateHeaders: false,
}

func Server(router *gin.Engine) error {
	s := gracehttp.Serve(
		&http.Server{
			Addr:           ":9090",
			Handler:        router,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		})
	return s
}
