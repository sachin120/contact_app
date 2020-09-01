package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp ...
func StartApp() {

	mapUrls()
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
