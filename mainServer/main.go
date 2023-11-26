package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user_growth/pb"
	"user_growth/ugServer"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/26 18:01
//

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	srv := grpc.NewServer()

	// 注册服务
	pb.RegisterUserCoinServer(srv, &ugServer.UgCoinServer{})
	pb.RegisterUserGradeServer(srv, &ugServer.UgGradeServer{})

	// 启动服务
	log.Printf("server listening ad : %v\n", lis.Addr())
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
