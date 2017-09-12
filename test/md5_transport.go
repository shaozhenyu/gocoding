package md5transport

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Transport struct {
	http.RoundTripper
}

func NewTransport(transport http.RoundTripper) http.RoundTripper {
	return &Transport{transport}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("aaaaa")
	if req == nil || req.Body == nil {
		return t.RoundTripper.RoundTrip(req)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println("body:", string(body))

	if len(body) > 0 {
		bodyMd5 := md5.Sum(body)
		req.Header.Add("X-Md5", hex.EncodeToString(bodyMd5[:]))
	}

	return t.RoundTripper.RoundTrip(req)
}
