package utils

import (
	"io"
	"log/slog"
	"net/http"
)

// 定义一个请求网址的通用方法
func NewReq(method, reqUrl string, reqBody io.Reader, reqHeaderMap map[string]string) []byte {

	client := http.Client{}

	req, err := http.NewRequest(method, reqUrl, reqBody)
	if err != nil {
		slog.Error("创建请求出错", "error", err)
		return nil
	}

	for key, value := range reqHeaderMap {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		slog.Error("请求错误", "error", err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		slog.Error("请求响应出错", "HttpMsg", res.Status, "HttpCode", res.StatusCode)
		return nil
	}

	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("读取返回内容出错", "error", err)
		return nil
	}

	return bodyByte
}
