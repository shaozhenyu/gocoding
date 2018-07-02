package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			resp, err := http.PostForm("http://uat-up-profit.bilibili.co/allowance/api/x/internal/growup/up/withdraw/success", url.Values{"order_no": {"1098"}, "trade_status": {"2"}})
			if err != nil {
				fmt.Println(err)
				return
			}
			resp.Body.Close()
		}()
	}
	fmt.Println("ok")
	time.Sleep(100 * time.Second)
	fmt.Println("ok1")
}
