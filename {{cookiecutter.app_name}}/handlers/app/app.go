package app

import (
	"github.com/gin-gonic/gin"
)

func AppRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "This is root",
	})
}
