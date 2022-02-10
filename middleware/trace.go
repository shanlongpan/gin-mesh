/**
* @Author:Tristan
* @Date: 2021/12/1 2:22 下午
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shanlongpan/gin-mesh/consts"
)

func NewTrackingId() string {
	return uuid.New().String()
}

// Gin Middleware with default header
func TrackingId() gin.HandlerFunc {
	return TrackingIdWithCustomizedHeader(consts.DefaultTraceIdHeader)
}

// Gin Middleware with cusomized header
func TrackingIdWithCustomizedHeader(head string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tId := c.GetHeader(head)
		// Generate TrackingID if not exist
		if tId == "" {
			tId = NewTrackingId()
			c.Header(head, tId)
		}

		// Set in Context
		c.Set(head, tId)
		c.Next()
	}
}
