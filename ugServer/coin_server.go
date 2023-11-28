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
		return nil, status.Errorf(codes.Internal, err.Error())
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
	coinUserSvc := service.NewCoinUserService(ctx)
	uid := int(in.Uid)
	data, err := coinUserSvc.GetByUid(uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	d := models.CoinUserToMessage(data)
	out := &pb.UserCoinInfoReply{
		Data: d,
	}
	return out, nil
}

// UserDetails 获取用户的积分明细列表
func (u *UgCoinServer) UserDetails(ctx context.Context, in *pb.UserDetailsRequest) (*pb.UserDetailsReply, error) {
	uid := int(in.Uid)
	page := int(in.Page)
	size := int(in.Size)
	coinDetailSvc := service.NewCoinDetailService(ctx)
	datalist, total, err := coinDetailSvc.FindByUid(uid, page, size)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	dlist := make([]*pb.TbCoinDetail, len(datalist))
	for i := range datalist {
		dlist[i] = models.CoinDetailToMessage(&datalist[i])
	}
	out := &pb.UserDetailsReply{
		Datalist: dlist,
		Total:    int32(total),
	}
	return out, nil
}

// UserCoinChange 调整用户积分-奖励和惩罚都是这个接口
func (u *UgCoinServer) UserCoinChange(ctx context.Context, in *pb.UserCoinChangeRequest) (*pb.UserCoinChangeReply, error) {
	uid := int(in.Uid)
	task := in.Task
	coin := int(in.Coin)
	taskInfo, err := service.NewCoinTaskService(ctx).GetByTask(task)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	//if taskInfo == nil {
	//	return nil, status.Errorf(codes.Internal, errors.New("r"))
	//}
	// 插入详情
	coinDetail := models.TbCoinDetail{
		Uid:    uid,
		TaskId: int(taskInfo.Id),
		Coin:   coin,
	}
	err = service.NewCoinDetailService(ctx).Save(&coinDetail)
	if err != nil {
		return nil, err
	}
	// 更新用户信息
	coinUserSvc := service.NewCoinUserService(ctx)
	coinUser, err := coinUserSvc.GetByUid(uid)
	if err != nil {
		return nil, err
	}
	if coinUser == nil {
		coinUser = &models.TbCoinUser{
			Uid:   uid,
			Coins: coin,
		}
	} else {
		coinUser.Coins += coin
		coinUser.SysCreated = nil
		coinUser.SysUpdated = nil
	}

	err = coinUserSvc.Save(coinUser)
	if err != nil {
		return nil, err
	}
	out := &pb.UserCoinChangeReply{
		User: models.CoinUserToMessage(coinUser),
	}
	return out, nil
}
