package activity_handler

import (
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/data_store/Splunk"
	"bitbucket.org/makeusmobile/makeus-golang-framework/src/models"
	. "bitbucket.org/makeusmobile/makeus-golang-framework/src/utility"

	"encoding/json"
	"net/http"
	// "time"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
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
