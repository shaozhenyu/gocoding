package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UpInfo struct {
	Id       int64 `json:"id"`
	Mid      int64 `json:"mid"`
	SignedAt int64 `json:"signed_at"`
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/api/x/admin/growup/up/signed", handler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("11111111")
	// m := make([]UpInfo, 3)
	// m[0] = UpInfo{
	// 	Id:       1,
	// 	Mid:      212230,
	// 	SignedAt: 1515479252,
	// }
	// m[1] = UpInfo{
	// 	Id:       2,
	// 	Mid:      2019740,
	// 	SignedAt: 1515480544,
	// }
	// m[2] = UpInfo{
	// 	Id:       3,
	// 	Mid:      275981,
	// 	SignedAt: 1515480642,
	// }
	m := UpInfo{
		Id:       1,
		Mid:      212230,
		SignedAt: 1515479252,
	}
	ret, _ := json.Marshal(&m)
	fmt.Fprintf(w, string(ret))
}
