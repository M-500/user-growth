package ugServer

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_growth/pb"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/26 18:02
//

type UgCoinServer struct {
	pb.UnimplementedUserCoinServer
}

func (u *UgCoinServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksReply, error) {
	//coinTaskSvc := service.NewCoinTaskService(ctx)
	//coinTaskSvc.FindAllPager()
	return nil, nil
}
func (u *UgCoinServer) UserCoinInfo(ctx context.Context, in *pb.UserCoinInfoRequest) (*pb.UserCoinInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCoinInfo not implemented")
}
func (u *UgCoinServer) UserDetails(ctx context.Context, in *pb.UserDetailsRequest) (*pb.UserDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetails not implemented")
}
func (u *UgCoinServer) UserCoinChange(ctx context.Context, in *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCoinChange not implemented")
}
