package characterutil

import (
	"encoding/json"
	"strings"
)

// StitchingBuilderStr 字符串拼接
func StitchingBuilderStr(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}

// IsJson 判断字符串是否为json字符串
func IsJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
