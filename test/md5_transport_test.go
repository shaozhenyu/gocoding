package md5transport

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	ast := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md5Str := r.Header.Get("X-Md5")
		b, err := ioutil.ReadAll(r.Body)
		ast.Nil(err)
		if len(b) == 0 {
			ast.Equal("", md5Str, r.URL.Path)
			return
		}

		fmt.Println("bbbbb")

		hexB := md5.Sum(b)
		md5exp := hex.EncodeToString(hexB[:])
		ast.Equal(md5exp, md5Str, r.URL.Path)
		fmt.Println(md5exp, md5Str)
		return
	}))
	defer ts.Close()

	md5tp := NewTransport(http.DefaultTransport)
	client := http.Client{Transport: md5tp}

	_, err := client.Get(ts.URL + "/testget")
	ast.Nil(err)

	for n := 1; n <= 5000; n *= 2 {
		path := "/" + strconv.Itoa(n)
		b := make([]byte, n)
		rand.Read(b)
		reader := bytes.NewBuffer(b)

		//fmt.Println("reader:", reader)
		_, err := client.Post(ts.URL+path, "application/octet-stream", reader)
		//fmt.Println("error:", err)
		ast.Nil(err)
	}
}
