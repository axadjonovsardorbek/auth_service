package service

import (
	m "auth-service/models"
	"fmt"
	"log"
	"net/rpc"
)

type UserService struct {
	Client *rpc.Client
}

func NewUserService(client *rpc.Client) *UserService {
	return &UserService{Client: client}
}

func (u *UserService) Register(req m.RegisterReq) (*m.User, error) {
	res := new(m.User)

	err := u.Client.Call("User.Register", req, &res)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetUser(id int) (*m.UserRes, error) {
	res := new(m.UserRes)

	err := u.Client.Call("User.GetUserData", id, &res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func (u *UserService) Login(req *m.LoginReq) bool {
	res := new(m.LoginRes)

	fmt.Println(req)

	err := u.Client.Call("User.Login", req, &res)
	if err != nil {
		log.Println(err)
		return false
	}

	return res.Correct
}
