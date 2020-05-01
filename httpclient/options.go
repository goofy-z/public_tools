package httpclient

import "time"

type Option func(*Client)

// 设置超时时间
func SetTimeout(timeout time.Duration) Option{
	return func(client *Client) {
		client.timeout = timeout
	}
}

// 设置重试次数
func SetRetry(count int) Option{
	return func(client *Client) {
		client.retry = count
	}
}

// 设置代理
func SetProxy(proxyUrl string) Option{
	return func(client *Client) {
		client.proxy = true
		client.proxyUrl = proxyUrl
	}
}

