package extend

import (
	"crypto/md5"
	"encoding/hex"
)

// 返回一个32位md5加密后的字符串
func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str))
    return hex.EncodeToString(h.Sum(nil))
}

// 16位md5加密后字符串
func Md5To16(str string) string {
    return Md5(str)[8:24]
}