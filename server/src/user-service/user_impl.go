package user_service

import (
	"context"
	"fmt"
)

type UserServiceServerImpl struct{}
func (s UserServiceServerImpl) QueryUser(ctx context.Context, user *RequestUser) (*ResponseUser, error){
	fmt.Println(ctx)
	fmt.Println(user)
	responseUser := &ResponseUser{Username:"张三，你好："+user.Uid,Age:11,Telphones:[]string {"11111","2222"},}
	return responseUser,nil
}
func (s UserServiceServerImpl) StreamUser( u UserService_StreamUserServer) error {
	user, err := u.Recv()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
		u.Send(&ResponseUser{Username:"李四，你好："+user.Uid,Age:77,Telphones:[]string {"aaaa","bbbb"},})
	}
	return nil
}


