package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"jtthink/src/Users"
	"log"
)

type UserService struct{}
func (g *UserService) Test(ctx context.Context, req *Users.UserRequest, rsp *Users.UserResponse) error {
	rsp.Ret="users"+req.Id
	return nil
}
func NewUserService() *UserService  {
	return &UserService{}
}
func main()  {
//api.jtthink.com --- prefix
	service:=micro.NewService(
		micro.Name("api.jtthink.com.user"))
	service.Init()

	err:=Users.RegisterUserServiceHandler(service.Server(),NewUserService())
	if err!=nil{
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}



}
