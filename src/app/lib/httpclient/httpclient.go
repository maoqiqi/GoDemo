package httpclient

import (
	"net/http"
	"time"
	"io"
	"fmt"
	"strconv"
	"strings"
)

const (
	FORM_URLENCODED = "application/x-www-form-urlencoded"
	JSON            = "application/json"
	XML             = "application/xml"
)

type HttpClient struct {
	*http.Client
}

func Client(options ... func(*HttpClient)) *HttpClient {
	httpClient := &HttpClient{&http.Client{}}
	// 设置超时时间
	httpClient.Client.Timeout = 5 * time.Second
	// 应用参数
	for _, option := range options {
		option(httpClient)
	}
	return httpClient
}

// POST 请求
func (h *HttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return h.Do(req)
}

// GET 请求
func (h *HttpClient) Get(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, err
	}
	return h.Do(req)
}

func (h *HttpClient) Do(req *http.Request) (*http.Response, error) {
	startTime := time.Now()
	resp, err := h.Client.Do(req)
	statistics(startTime, req.URL.String(), resp, err)
	return resp, err
}

// 统计请求时间
func statistics(startTime time.Time, urlStr string, resp *http.Response, err error) {
	endTime := time.Now()
	diffTimeStr := fmt.Sprintf("%.6f", endTime.Sub(startTime).Seconds())
	diffTime, _ := strconv.ParseFloat(diffTimeStr, 64)
	fmt.Println(urlStr, "接口效率", FormatTime(diffTime))

	if err == nil {
		if resp.StatusCode == 200 {
			fmt.Println(urlStr, "成功")
		}
		fmt.Println(urlStr, "请求状态码", resp.StatusCode)
	} else {
		fmt.Println(urlStr, "请求错误", parseErr(err.Error()))
		if resp != nil {
			fmt.Println(urlStr, "请求状态码", resp.StatusCode)
		}
	}
}

func FormatTime(diffTime float64) string {
	var info string
	if diffTime < 0.05 {
		info = "0.00s到0.05s"
	} else if diffTime < 0.1 {
		info = "0.05s到0.1s"
	} else if diffTime < 0.5 {
		info = "0.1s到0.5s"
	} else if diffTime < 1 {
		info = "0.5s到1s"
	} else if diffTime < 2 {
		info = "1s到2s"
	} else if diffTime < 3 {
		info = "2s到3s"
	} else if diffTime < 4 {
		info = "3s到4s"
	} else if diffTime < 5 {
		info = "4s到5s"
	} else if diffTime < 10 {
		info = "5s到10s"
	} else {
		info = "10s到∞秒"
	}
	return info
}

func parseErr(str string) string {
	errs := make(map[string]string)
	errs["Client.Timeout"] = "HTTP请求超时(28)"
	errs["no such host"] = "DNS解析失败(6)"
	errs["unsupported protocol scheme"] = "网址格式不正确(3)"

	for k, v := range errs {
		if strings.Index(str, k) != -1 {
			return v
		}
	}

	return "错误"
}
