package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R403 forbidden
func R403(c *gin.Context, data interface{}, key string) {
	if key == "" {
		key = CommonForbidden
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	sendResponse(c, http.StatusForbidden, false, data, localeData.Message, localeData.Code)
}
