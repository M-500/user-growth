package common

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 15:43
//

func Now() *time.Time {
	now := time.Now()
	return &now
}
