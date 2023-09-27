package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R400 bad request
func R400(c *gin.Context, data interface{}, key string)  {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonBadRequest
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	 sendResponse(c, http.StatusBadRequest, false, data, localeData.Message, localeData.Code)
}
