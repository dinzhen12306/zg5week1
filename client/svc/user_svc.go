package svc

import (
	"google.golang.org/grpc"
	user "week1/server/userrpc"
)

type UserSvc struct {
}

func (s *UserSvc) UserRpcConn(port string) (user.UserClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return user.NewUserClient(conn), nil
}
