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

func MessageToCoinTask(data *pb.TbCoinTask) *TbCoinTask {
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
