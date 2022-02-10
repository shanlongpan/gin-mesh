package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanlongpan/gin-mesh/lib"
	"github.com/shanlongpan/gin-mesh/xlog"
	"github.com/shanlongpan/grpc-mesh/grpc/user"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/xds"
	"math/rand"
	"net/http"
	"time"
)

var UserService user.UserServiceClient

func init() {
	ctx := context.TODO()
	//target := fmt.Sprintf("xds:///grpc-mesh.echo-grpc.svc.cluster.local:%s", config.Conf.GrpcPort)
	target := "xds:///grpc-mesh.echo-grpc.svc.cluster.local:7070"
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	xlog.Println(ctx, "target", target)
	if err != nil {
		xlog.Errorln(ctx, err)
	}
	UserService = user.NewUserServiceClient(conn)
}
func create(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)

	nctx := lib.GetNewCtx(ctx)
	resp, err := UserService.Create(nctx, &user.CreateUserReq{User: &user.User{
		Name:  fmt.Sprintf("xiaoming_%d", d),
		Email: fmt.Sprintf("xiaoming_%d@163.com", d),
	}})

	if err != nil {
		xlog.Errorln(ctx, err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   "1111111",
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": resp.String(),
		})
	}
}

func read(ctx *gin.Context) {


	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	nctx := lib.GetNewCtx(ctx)
	res, err := UserService.Read(nctx, &user.ReadUserReq{
		Id: int64(d),
	})

	if err != nil {
		xlog.Errorln(ctx, err)
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
			"msg": "222222",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res.String(),
			"err": "444444",
		})
	}
}

func ping(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": d,
	})
}
