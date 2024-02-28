package api

import (
	"context"
	"week1/server/server"
	user "week1/server/userrpc"
)

type UserRpcServer struct {
	user.UnimplementedUserServer
}

func (u UserRpcServer) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	userInfo, err := server.CreateUser(req.User)
	if err != nil {
		return nil, err
	}
	return &user.CreateUserResp{
		User: userInfo,
	}, nil
}

func (u UserRpcServer) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	err := server.DeleteUser(req.UserID)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResp{}, nil
}

func (u UserRpcServer) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	updateUser, err := server.UpdateUser(req.User)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResp{
		User: updateUser,
	}, nil
}

func (u UserRpcServer) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	getUser, err := server.GetUser(req.Where)
	if err != nil {
		return nil, err
	}
	return &user.GetUserResp{
		User: getUser,
	}, nil
}

func (u UserRpcServer) GetUsers(ctx context.Context, req *user.GetUsersReq) (*user.GetUsersResp, error) {
	users, i, err := server.GetUsers(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	return &user.GetUsersResp{
		Users: users,
		Count: i,
	}, nil
}
