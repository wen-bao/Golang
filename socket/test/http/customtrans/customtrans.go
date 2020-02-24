package main

import (
	"fmt"
	"net/http"
)

// OurCustomTransport ...
type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// RoundTrip ...
func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情
	// 发起HTTP请求
	// 添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}

// Client ...
func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func main() {
	t := &OurCustomTransport{
		// ...
	}

	c := t.Client()
	resp, err := c.Get("http://www.oabnew.com")
	fmt.Println(resp, err)
	// ...
}
