/**
* @Author:Tristan
* @Date: 2021/12/31 5:34 下午
 */

package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/gin-mesh/consts"
)

func GetNewCtx(ctx *gin.Context) context.Context {
	nctx := context.Background()
	traceId, _ := ctx.Get(consts.DefaultTraceIdHeader)
	nctx = context.WithValue(nctx, consts.DefaultTraceIdHeader, traceId)
	return nctx
}
