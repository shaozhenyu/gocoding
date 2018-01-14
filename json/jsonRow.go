package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var result interface{}
	result = `{"aa": "aa", "bb": "bb"}`
	m := map[string]interface{}{
		"code":   0,
		"msg":    "ok",
		"result": result,
	}

	str, ok := result.(string)
	if ok {
		rowData := json.RawMessage(str)
		m["result"] = rowData
	}

	b, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("body: ", string(b))
}
