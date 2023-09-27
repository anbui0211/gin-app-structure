package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R404 not found
func R404(c *gin.Context, data interface{}, key string) {
	if key == "" {
		key = CommonNotFound
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	sendResponse(c, http.StatusNotFound, false, data, localeData.Message, localeData.Code)
}
