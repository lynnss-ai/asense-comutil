package netutil

import "net/url"

// UrlEncode URL编码
// apiUrl: 需要编码的URL
// params: URL参数
//
// return: 编码后的URL
func UrlEncode(apiUrl string, params map[string]string) (string, error) {
	_url, err := url.Parse(apiUrl)
	if err != nil {
		return "", err
	}
	query := _url.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	_url.RawQuery = query.Encode()
	r, _ := url.QueryUnescape(_url.String())
	return r, nil
}
