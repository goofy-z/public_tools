package httpclient

import (
	"bytes"
	"fmt"
	"git.code.oa.com/infosec_basicPlatform_devops/devops_sdk/zlog"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	// 实际client对象
	httpclient *http.Client

	// 超时时间单位秒
	timeout time.Duration

	// 重试次数
	retry int

	// 是否使用代理
	proxy bool
	proxyUrl string

}

type Httper interface {
	Do(*http.Request) (*http.Response, error)
}


func NewClient(opts ...Option)*Client{
	client := Client{
		timeout:    3, // 秒为单位
		retry:      0, // 重试0次
		proxy:      false,
		proxyUrl:   "",
	}
	// 加载额外配置
	for _, opt := range opts {
		opt(&client)
	}
	// 初始化http.Client对象, client
	if client.httpclient == nil {
		client.httpclient = &http.Client{
			Timeout: client.timeout * time.Second,
		}
	}
	// 设置proxy
	if client.proxy{
		proxyURL, _ := url.Parse(client.proxyUrl)
		trans := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		client.httpclient.Transport = trans
	}
	return &client

}

func(c *Client)Do(req *http.Request)(*http.Response, error){
	req.Close = true
	var response *http.Response
	var bodyReader *bytes.Reader

	// 重新复制req的body，避免重试时因为读过一次之后将reader对象关闭
	if req.Body != nil {
		reqData, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(reqData)
		// 因为req.body是ReadCloser对象需要改造reader对象
		req.Body = ioutil.NopCloser(bodyReader)
	}
	var err error
	for i := 0; i <= c.retry; i++ {
		if response != nil {
			response.Body.Close()
		}
		response, err = c.httpclient.Do(req)
		// 需要将Reader的探针重置到起点
		if bodyReader != nil{
			_, _ = bodyReader.Seek(0, 0)
		}
		if err != nil{
			zlog.Errorf("请求失败, 原因: %s", err.Error())
			continue
		}
		if response.StatusCode != http.StatusOK{
			err = fmt.Errorf("http statusCode is not 200")
			zlog.Errorf("请求状态错误")
			continue
		}
	}
	return response, err
}
