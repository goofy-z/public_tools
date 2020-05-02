package httpclient

import (
	"github.com/qq1141000259/public_tools/zlog"
	"io"
	"net/http"
)

func(c *Client)GET(url string, header http.Header)(*http.Response, error){
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil{
		zlog.Errorf("构造[GET]请求失败, url: %s, 原因: %s", url, err.Error())
		return nil, err
	}
	req.Header = header
	return c.Do(req)
}

func(c *Client)POST(url string, header http.Header, body io.Reader)(*http.Response, error){
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil{
		zlog.Errorf("构造[POST]请求失败, url: %s, 原因: %s", url, err.Error())
		return nil, err
	}
	req.Header = header
	return c.Do(req)
}

func(c *Client)PUT(url string, header http.Header, body io.Reader)(*http.Response, error){
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil{
		zlog.Errorf("请求[PUT]失败, url: %s, 原因: %s", url, err.Error())
		return nil, err
	}
	req.Header = header
	return c.Do(req)
}

func(c *Client)DELETE(url string, header http.Header)(*http.Response, error){
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil{
		zlog.Errorf("请求[DELETE]失败, url: %s, 原因: %s", url, err.Error())
		return nil, err
	}
	req.Header = header
	return c.Do(req)
}