package mongodb

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/config"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/utility/error"

	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// MongoDB is struct for singleton
type MongoDB struct {
	Session *mgo.Session // Session
}

// Global instance
var Instance = MongoDB{Session: nil}

func InitMongoDB() *mgo.Session {

	mongodb := config.MONGODB_ADDR
	session, err := mgo.Dial(mongodb)

	if err != nil {
		log.panic(error.ErrNotFountInstant)
	}

	Instance.Session = session

	return session
}

func connector(next gin.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {

		if Instance.Session != nil {
			log.panic(error.ErrNotFountInstant)
		}

		s := Instance.Session.Clone()
		defer s.Close()
		c.Set("mongodb", s)

		next(c)
	})
}
