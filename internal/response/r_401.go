package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R401 unauthorized
func R401(c *gin.Context, data interface{}, key string)  {
	if key == "" {
		key = CommonUnauthorized
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	 sendResponse(c, http.StatusUnauthorized, false, data, localeData.Message, localeData.Code)
}
