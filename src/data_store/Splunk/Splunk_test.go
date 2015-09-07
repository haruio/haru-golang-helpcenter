package Splunk

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	router := gin.New()

	// router.addRoute("GET", "/singapore/screen", HandlersChain{func(_ *Context) {}})

	// assert.Len(t, router.trees, 1)
	// assert.NotNil(t, router.trees.get("GET"))
	// assert.Nil(t, router.trees.get("POST"))

}
