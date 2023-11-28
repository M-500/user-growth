package ugServer

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user_growth/models"
	"user_growth/pb"
	"user_growth/service"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/26 18:02
//

type UgCoinServer struct {
	pb.UnimplementedUserCoinServer
}

// ListTasks 获取所有的积分任务列表
func (u *UgCoinServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksReply, error) {
	coinTaskSvc := service.NewCoinTaskService(ctx)
	dataList, err := coinTaskSvc.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]*pb.TbCoinTask, len(dataList))
	for i, task := range dataList {
		result[i] = models.CoinTaskToMessage(&task)
	}
	out := &pb.ListTasksReply{Datalist: result}
	return out, nil
}

// UserCoinInfo 获取用户的积分信息
func (u *UgCoinServer) UserCoinInfo(ctx context.Context, in *pb.UserCoinInfoRequest) (*pb.UserCoinInfoReply, error) {
	//uid := int(in.Uid)
	//coinTaskSvc := service.New(ctx)
	//data,err := coinTaskSvc.
	return nil, status.Errorf(codes.Unimplemented, "method UserCoinInfo not implemented")
}

// UserDetails 获取用户的积分明细列表
func (u *UgCoinServer) UserDetails(ctx context.Context, in *pb.UserDetailsRequest) (*pb.UserDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDetails not implemented")
}

// UserCoinChange 调整用户积分-奖励和惩罚都是这个接口
func (u *UgCoinServer) UserCoinChange(ctx context.Context, in *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCoinChange not implemented")
}
