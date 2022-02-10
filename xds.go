/**
* @Author:Tristan
* @Date: 2022/2/10 7:46 下午
 */

package main

import (
	"fmt"
	"github.com/shanlongpan/grpc-mesh/grpc/user"
	"google.golang.org/grpc"
	"context"
	"math/rand"
	"time"
)

var UserService user.UserServiceClient

func init() {
	ctx := context.TODO()
	//target := fmt.Sprintf("xds:///grpc-mesh.echo-grpc.svc.cluster.local:%s", config.Conf.GrpcPort)
	target := "xds:///grpc-mesh.echo-grpc.svc.cluster.local:7070"
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())

	if err != nil {
		fmt.Println("init",err.Error())
	}
	UserService = user.NewUserServiceClient(conn)
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	d := rand.Intn(2000)
	resp, err := UserService.Create(context.TODO(), &user.CreateUserReq{User: &user.User{
		Name:  fmt.Sprintf("xiaoming_%d", d),
		Email: fmt.Sprintf("xiaoming_%d@163.com", d),
	}})
	fmt.Println(err)
	fmt.Println(resp.String())
}