package handler

import "github.com/gin-gonic/gin"

type Common struct{}

// Ping godoc
// @tags Common
// @summary Ping
// @id ping
// @security ApiKeyAuth
// @accept json
// @produce json
// @router /ping [get]
func (Common) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
