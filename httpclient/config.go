package httpclient

import (
	"time"
)

var (
	// 设置默认的超时时间
	DefaultTimeout = time.Duration(3) * time.Second

	// 设置默认的重试次数
	DefaultRetry = 0
)
