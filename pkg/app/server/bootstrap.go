package server

import (
	"igapp/pkg/app/route"
	"igapp/pkg/app/server/initialize"
	"log"

	"github.com/gin-gonic/gin"
)

// Bootstrap ...
func Bootstrap(r *gin.Engine) {
	log.Println("Initializing server ...")

	initialize.Init(r)

	route.Init(r)
}
