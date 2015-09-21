package couchbase

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config/error"

	"log"

	"github.com/couchbaselabs/gocb"
	"github.com/couchbaselabs/gocb/gocbcore"
	"github.com/gin-gonic/gin"
)

// Couchbase is struct for singleton
type Couchbase struct {
	Bucket *gocb.Bucket // Bucket
}

// Global instance
var Instance = &Couchbase{Bucket: nil}

func InitCouchbase() *gocb.Bucket {

	myCluster, err := gocb.Connect(config.COUCHBASE_ADDR)
	if err != nil {
		log.Panic(error.ErrNotFountInstant)
	}

	myBucket, err := myCluster.OpenBucket(config.COUCHBASE_BUCKETNAME, config.COUCHBASE_PASSWORD)
	if err != nil {
		log.Panic(error.ErrNotFountInstant)
	}

	Instance.Bucket = myBucket

	SetLogger(LogStdOutLogger()) //couchbase internal package enable logger("log" package)

	return myBucket
}

func Connector(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		if Instance.Bucket != nil {
			log.Panic(error.ErrNotFountInstant)
		}

		c.Set("couchbase", Instance.Bucket)

		next(c)
	})
}

func SetLogger(logger gocbcore.Logger) {
	gocbcore.SetLogger(logger)
}
