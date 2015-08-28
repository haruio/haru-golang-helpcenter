package handlers

import (
	"../database/Splunk"
	//"../models"
	. "../utility"

	"encoding/json"
<<<<<<< HEAD
	// "log"
=======
>>>>>>> 181072279ce089efcb68141fcd9205f388961e7d
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////              AppContext              ///////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type AppContext struct {
	Splk *Splunk.Splunk
}

func (this *AppContext) Request(c *gin.Context) {

	var v interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&v)
	GinErrCheck(err, c)

	body, err := json.Marshal(v)
	GinErrCheck(err, c)

	_, err = this.Splk.Connection.Write(body)
	GinErrCheck(err, c)

	c.JSON(http.StatusOK, gin.H{"screen": v})
}

func (this *AppContext) Close() {
	this.Splk.Close()
}
