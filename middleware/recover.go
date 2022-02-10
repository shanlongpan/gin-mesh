/**
* @Author:Tristan
* @Date: 2021/12/1 11:45 上午
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/gin-mesh/xlog"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			msg := fmt.Sprintf("message %v stack: %+v", errorToString(r), string(debug.Stack()))
			xlog.Errorln(c,msg)

			if gin.Mode() == gin.ReleaseMode {
				msg = http.StatusText(http.StatusInternalServerError)
			}
			// abort 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  msg,
				"data": nil,
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
