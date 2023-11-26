package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
	"user_growth/pb"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/26 18:01
//

func main() {
	add := flag.String("addr", "localhost:80", "the addr to connect to ")
	conn, err := grpc.Dial(*add)
	if err != nil {
		log.Fatalf("faile to connect : %v\n", err)
	}
	defer conn.Close()
	// 请求服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 新建客户端
	cCoin := pb.NewUserCoinClient(conn)
	cGrade := pb.NewUserGradeClient(conn)
	// 测试
	r1, err1 := cCoin.ListTasks(ctx, &pb.ListTasksRequest{})
	if err1 != nil {
		log.Printf("%v", err1)
	} else {
		log.Printf("%+v\n", r1.GetDatalist())
	}
	r2, err2 := cGrade.ListGrades(ctx, &pb.ListGradesRequest{})
	if err2 != nil {
		log.Printf("%v", err2)
	} else {
		log.Printf("%+v\n", r2.GetDatalist())
	}
}
