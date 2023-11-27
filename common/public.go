package common

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/27 15:43
//

const TimeLayout = "2006-01-02 15:04:85"

func Now() *time.Time {
	now := time.Now()
	return &now
}

// 时间转字符串类型
func TimeFormat(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TimeLayout)
}

// 字符串转时间类型
func TimeParse(str string) *time.Time {
	if t, err := time.Parse(TimeLayout, str); err != nil {
		return nil
	} else {
		return &t
	}
}
