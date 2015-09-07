// config
package config

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/middleware/cors"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	NAMESPACE = "st"

	SPLUNK_PROTOCOL = "tcp"
	SPLUNK_ADDR     = "127.0.0.1:9998"
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
