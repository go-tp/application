package extend

import "time"

// 将时间戳转换成字符串日期
func Date(tn int) string {
	timeTemplate := "2006-01-02 15:04:05"
	t := time.Unix(int64(tn),0).Format(timeTemplate)
	return t
}

// 返回当前时间戳
func Now() int64 {
	return time.Now().Unix()
}