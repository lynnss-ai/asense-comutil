package base64util

import "encoding/base64"

// Base64UrlEncode URL base64编码
// data: 需要编码的URL字符串
// return: 编码后的值
func Base64UrlEncode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

// Base64UrlDecode URL base64解码
// data: 需要解码的URL字符串
// return: 解码后的值
func Base64UrlDecode(data string) string {
	r, _ := base64.URLEncoding.DecodeString(data)
	return string(r)
}
