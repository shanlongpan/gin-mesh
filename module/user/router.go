package user

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/user")
	v1.GET("/create", create)
	v1.GET("/read", read)
	v1.GET("/ping", ping)

}
