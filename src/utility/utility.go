package utillity

import (
	"log"
	"net/http"

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
