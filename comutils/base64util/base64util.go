package base64util

import "encoding/base64"

// Base64Encode base64编码
// data: 需要编码的字符串
// return: 编码后的字符串
// Example: Base64Encode("hello") => "aGVsbG8="
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode base64解码
// data: 需要解码的字符串
// return: 解码后的字符串
// Example: Base64Decode("aGVsbG8=") => "hello"
func Base64Decode(data string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(data)
	return string(decodeBytes)
}
