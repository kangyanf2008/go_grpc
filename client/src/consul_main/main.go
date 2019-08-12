package main

import (
	"consul_client/grpclb"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
	"user-service"
)
const (
	address     = "172.21.48.3:8080"
)

var (
	//serv = flag.String("service", "hello_service", "service name")
	//address = flag.String("address", "127.0.0.1:8080", "listening port")
	//reg = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	conn, err := grpc.Dial(
		"",
		grpc.WithInsecure(),
		// 开启 grpc 中间件的重试功能
/*		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Duration(1)*time.Millisecond)), // 重试间隔时间
				grpc_retry.WithMax(3),                                             // 重试次数
				grpc_retry.WithPerRetryTimeout(time.Duration(5)*time.Millisecond), // 重试时间
				// 返回码为如下值时重试
				grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),*/

		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)), // 重试间隔时间
				grpc_retry.WithMax(3),                                             // 重试次数
				grpc_retry.WithPerRetryTimeout(time.Second), // 重试时间
				// 返回码为如下值时重试
				grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),
		// 负载均衡，使用 consul 作服务发现
		grpc.WithBalancer(grpc.RoundRobin(grpclb.NewConsulResolver(
			"172.21.48.3:8501", "grpc.health.v1.user-server",
		))),

/*		grpc.WithBalancer(grpc.RoundRobin(grpclb.NewPseudoResolver(
			[]string{
				"127.0.0.1:8080",
				"127.0.0.1:9080",
			},
		))),*/
	)
	if err != nil {
		fmt.Printf("dial failed. err: [%v]\n", err)
		return
	}
	defer conn.Close()
/*	//睡眠10S，由于grpc需要连接注册中心初始化数据，因此需要休息，不然会有调用时还没有进行初始化
	//time.Sleep(3*time.Second)*/
	var i int64 = 1
	for {
		//业务代码
		c := user_service.NewUserServiceClient(conn)
		//同步请求
		_, err := c.QueryUser(context.Background(), &user_service.RequestUser{Uid:"bbbbbbbbbbb"})
/*		fmt.Println(r.GetTelphones())
		fmt.Println(r.GetAge())
		fmt.Println(r.GetUsername())*/
		//fmt.Println(r)

/*		//异步 调用
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = c.QueryUser(ctx, &user_service.RequestUser{Uid:"bbbbbbbbbbb"})
		fmt.Println(r.GetTelphones())
		fmt.Println(r.GetAge())
		fmt.Println(r.GetUsername())*/

		if i %10000 == 0 {
			fmt.Println(fmt.Sprint(time.Now())+"第"+fmt.Sprint(i)+"次")
			time.Sleep(time.Second)
		}
		i++
		if err != nil {
			//出现异常打印日志会奔溃  log.Fatalf("could not greet: %v", err)
			fmt.Printf("could not greet: %v", err)
		}
	}


}