package main

import (
	"fmt"
	"net/http"
)

type TestRoundTriper struct {
	rt   http.RoundTripper
	test string
}

func (c *TestRoundTriper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("request : ", c.test)
	fmt.Println(req)
	resp, err := c.rt.RoundTrip(req)
	fmt.Println("response: ")
	fmt.Println(resp)
	return resp, err
}

func NewRoundTriper(rt http.RoundTripper, str string) *TestRoundTriper {
	return &TestRoundTriper{
		rt:   rt,
		test: str,
	}
}

func main() {
	client := &http.Client{
		Transport: NewRoundTriper(http.DefaultTransport, "szy"),
	}

	client.Get("http://www.baidu.com")
}
