package mongodb

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/utility/error"

	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	Session *mgo.Session // Session
}

var Instantiated *MongoDB = nil

func InitMongoDB() *mgo.Session {
	mongodb := config.MONGODB_ADDR
	session, err := mgo.Dial(mongodb)
	if err != nil {
		panic(error.ErrNotFountInstant)
	}
	Instantiated = session
	return session
}

func connector(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		if Instantiated.Session != nil {
			log.panic(error.ErrNotFountInstant)
		}

		s := Instantiated.Session.Clone()
		defer s.Close()
		c.Set("mongodb", s)

		next(c)
	})
}
