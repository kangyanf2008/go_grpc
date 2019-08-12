package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user_service.NewUserServiceClient(conn)

	var i int64 = 1
	for{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		//ctx:=context.Background()
		_, err := c.QueryUser(ctx, &user_service.RequestUser{Uid:"bbbbbbbbbbb"})
		if i %10000 == 0 {
			fmt.Println(fmt.Sprint(time.Now())+"第"+fmt.Sprint(i)+"次")
			time.Sleep(time.Second)
		}
		i++
		//fmt.Println(r.GetTelphones())
		//fmt.Println(r.GetAge())
		//fmt.Println(r.GetUsername())
		//ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
		//defer cancel2()
		//aa , _ := c.StreamUser(ctx2,)

		//aa , _ := c.StreamUser(ctx,)
		//ee := aa.Send(&user_service.RequestUser{Uid:"mmmmmmmmmmmmmmmmmmmmmm"})
		//if ee == nil {
		//	responseUser, ee:= aa.Recv()
		//	fmt.Println(ee)
		//	fmt.Println(responseUser.Username)
		//	fmt.Println(responseUser.Age)
		//	fmt.Println(responseUser.Telphones)
		//}
		if err != nil {
			//log.Fatalf("could not greet: %v", err)
		}
	}

}

type MyCallOption struct {

}
// returns a non-nil error, the RPC fails with that error.
func (MyCallOption) before(option *grpc.CallOption) error {
	return nil
}

// after is called after the call has completed.  after cannot return an
// error, so any failures should be reported via output parameters.
func (MyCallOption) after(option *grpc.CallOption) {

}