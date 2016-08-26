package main

import (
	"fmt"
	"encoding/json"
)

type reqBody struct {
	req string
}

func (r *reqBody) jsonToMap() (ret map[string]interface{}, err error) {
	if err := json.Unmarshal([]byte(r.req), &ret); err != nil {
		return nil, err
	}

	return ret, nil
}


func main() {
	var r reqBody
	r.req = `{"name":"szy", "age": 18, "gender":"man"}`
	if rMap, err := r.jsonToMap(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rMap["name"], rMap["age"], rMap["gender"])
	}
}
