package main

import (
	"consul_client"
	grpcsr "consul_client/grpcsr"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"user-service"
)

func main() {

	server :=grpc.NewServer()
	user_service.RegisterUserServiceServer(server, &user_service.UserServiceServerImpl{})
	grpc_health_v1.RegisterHealthServer(server,  &consul_client.HealthImpl{})

	//方式一
	go func() {
			e := grpcsr.NewConsulRegisterS("172.21.48.3:8501","user-server",8080).Register()
			if e != nil {
				fmt.Printf(e.Error())
			}
		}()


	// 方式二  从 consul 读取配置文件
/*	config := viper.New()
	config.AddRemoteProvider("consul", "172.21.48.3:8501", "D:\\kangyanf\\go_grpc\\server\\src\\configs\\user-service.json")
	config.SetConfigType("json")
	if err := config.ReadRemoteConfig(); err != nil {
		panic(err)
	}
	config.BindPFlags(pflag.CommandLine)

	// 使用 consul 注册服务
	register := consul_client.NewConsulRegisterS()
	config.Sub("register").Unmarshal(register)
	register.Port = 8500
	if err := register.Register(); err != nil {
		panic(err)
	}*/


	list , err := net.Listen("tcp",":8080")
	if err != nil{
		fmt.Println(err)
	}
	if err := server.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
