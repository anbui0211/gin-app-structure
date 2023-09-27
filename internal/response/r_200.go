package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// R200 response success
func R200(c *gin.Context, data interface{}, key string)  {
	// Get lang from echo context, if handle multilingualism

	if key == "" {
		key = CommonSuccess
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Message = key
	}
	 sendResponse(c, http.StatusOK, true, data, localeData.Message, localeData.Code)
}
