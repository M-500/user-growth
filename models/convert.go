//@Author: wulinlin
//@Description:
//@File:  convert
//@Version: 1.0.0
//@Date: 2023/11/27 20:08

package models

import (
	"user_growth/common"
	"user_growth/pb"
)

func CoinTaskToMessage(data *TbCoinTask) *pb.TbCoinTask {
	d := &pb.TbCoinTask{
		Id:        int32(data.Id),
		Task:      data.Task,
		Coin:      int32(data.Coin),
		Limit:     int32(data.Limit),
		SysStatus: int32(data.SysStatus),
	}
	d.Start = common.TimeFormat(data.Start)
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}

func CoinTaskToObject(data *pb.TbCoinTask) *TbCoinTask {
	d := &TbCoinTask{
		Id:    uint(data.Id),
		Task:  data.Task,
		Coin:  int(data.Coin),
		Limit: int(data.Limit),

		SysStatus: int(data.SysStatus),
	}
	d.Start = common.TimeParse(data.Start)
	d.SysCreated = common.TimeParse(data.Start)
	d.SysUpdated = common.TimeParse(data.Start)
	return d
}

func CoinUserToMessage(data *TbCoinUser) *pb.TbCoinUser {
	d := &pb.TbCoinUser{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		Coins:      int32(data.Coins),
		SysCreated: "",
		SysUpdated: "",
	}
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}
func CoinDetailToMessage(data *TbCoinDetail) *pb.TbCoinDetail {
	d := &pb.TbCoinDetail{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		TaskId:     int32(data.TaskId),
		Coin:       int32(data.Coin),
		SysCreated: "",
		SysUpdated: "",
	}
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}

// CoinDetailToObject message转model
func CoinDetailToObject(data *pb.TbCoinDetail) *TbCoinDetail {
	d := &TbCoinDetail{
		Id:         uint(data.Id),
		Uid:        int(data.Uid),
		TaskId:     int(data.TaskId),
		Coin:       int(data.Coin),
		SysCreated: nil,
		SysUpdated: nil,
	}
	d.SysCreated = common.TimeParse(data.SysCreated)
	d.SysUpdated = common.TimeParse(data.SysUpdated)
	return d
}

// GradeInfoToMessage model转message
func GradeInfoToMessage(data *TbGradeInfo) *pb.TbGradeInfo {
	d := &pb.TbGradeInfo{
		Id:          int32(data.Id),
		Title:       data.Title,
		Description: data.Description,
		Score:       int32(data.Score),
		Expired:     int32(data.Expired),
		SysCreated:  "",
		SysUpdated:  "",
	}
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}

// GradeInfoToObject message转model
func GradeInfoToObject(data *pb.TbGradeInfo) *TbGradeInfo {
	d := &TbGradeInfo{
		Id:          uint(data.Id),
		Title:       data.Title,
		Description: data.Description,
		Score:       int(data.Score),
		Expired:     int(data.Expired),
		SysCreated:  nil,
		SysUpdated:  nil,
	}
	d.SysCreated = common.TimeParse(data.SysCreated)
	d.SysUpdated = common.TimeParse(data.SysUpdated)
	return d
}

// GradePrivilegeToMessage model转message
func GradePrivilegeToMessage(data *TbGradePrivilege) *pb.TbGradePrivilege {
	d := &pb.TbGradePrivilege{
		Id:          int32(data.Id),
		GradeId:     int32(data.GradeId),
		Product:     data.Product,
		Function:    data.Function,
		Description: data.Description,
		Expired:     int32(data.Expired),
		SysCreated:  "",
		SysUpdated:  "",
		SysStatus:   int32(data.SysStatus),
	}
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}

// GradePrivilegeToObject message转model
func GradePrivilegeToObject(data *pb.TbGradePrivilege) *TbGradePrivilege {
	d := &TbGradePrivilege{
		Id:          uint(data.Id),
		GradeId:     int(data.GradeId),
		Product:     data.Product,
		Function:    data.Function,
		Description: data.Description,
		Expired:     int(data.Expired),
		SysCreated:  nil,
		SysUpdated:  nil,
		SysStatus:   int(data.SysStatus),
	}
	d.SysCreated = common.TimeParse(data.SysCreated)
	d.SysUpdated = common.TimeParse(data.SysUpdated)
	return d
}

// GradeUserToMessage model转message
func GradeUserToMessage(data *TbGradeUser) *pb.TbGradeUser {
	d := &pb.TbGradeUser{
		Id:         int32(data.Id),
		Uid:        int32(data.Uid),
		GradeId:    int32(data.GradeId),
		Expired:    "",
		Score:      int32(data.Score),
		SysCreated: "",
		SysUpdated: "",
	}
	d.Expired = common.TimeFormat(data.Expired)
	d.SysCreated = common.TimeFormat(data.SysCreated)
	d.SysUpdated = common.TimeFormat(data.SysUpdated)
	return d
}

// GradeUserToObject message转model
func GradeUserToObject(data *pb.TbGradeUser) *TbGradeUser {
	d := &TbGradeUser{
		Id:         uint(data.Id),
		Uid:        int(data.Uid),
		GradeId:    int(data.GradeId),
		Expired:    nil,
		Score:      int(data.Score),
		SysCreated: nil,
		SysUpdated: nil,
	}
	d.Expired = common.TimeParse(data.Expired)
	d.SysCreated = common.TimeParse(data.SysCreated)
	d.SysUpdated = common.TimeParse(data.SysUpdated)
	return d
}
