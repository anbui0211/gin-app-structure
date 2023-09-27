package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func sendResponse(c *gin.Context, httpCode int, success bool, data interface{}, message string, code int) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(httpCode, gin.H{
		"success": success,
		"data":    data,
		"message": message,
		"code":    code,
	})
}

// InvalidChecksum ...
func InvalidChecksum(c *gin.Context) {
	key := CommonInvalidChecksum
	localeData := GetByKey(key)
	sendResponse(c, http.StatusBadRequest, false, gin.H{}, localeData.Message, localeData.Code)
}

func RouteValidation(c *gin.Context, err error) {
	key := getMessage(err)

	// Return
	R400(c, nil, key)
}

func getMessage(err error) string {
	err1, ok := err.(validation.Errors)
	if !ok {
		err2, ok := err.(validation.ErrorObject)
		if ok {
			return err2.Message()
		}
		return err.Error()
	}
	for _, item := range err1 {
		if item == nil {
			continue
		}
		return getMessage(item)
	}
	return err.Error()
}
