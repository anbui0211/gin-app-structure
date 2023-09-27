package route

import (
	middlewareinternal "igapp/internal/util/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// Middleware ...
	r.Use(middlewareinternal.CORSMiddleware())

	v1 := r.Group("/api/v1")
	common(v1)
}
