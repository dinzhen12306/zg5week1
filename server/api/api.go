package api

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	user "week1/server/userrpc"
)

// 开启rpc服务
func OpenGrpcServer(port int, fun func(s *grpc.Server), tls credentials.TransportCredentials) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	var grpcServer *grpc.Server
	if tls == nil {
		grpcServer = grpc.NewServer()
	} else {
		grpcServer = grpc.NewServer(grpc.Creds(tls))
	}
	fun(grpcServer)
	reflection.Register(grpcServer)
	log.Println("rpc server addr: ", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}

// 注册服务
func RpcServerRegister(s *grpc.Server) {
	user.RegisterUserServer(s, new(UserRpcServer))
}
