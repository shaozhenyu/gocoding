package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Password int    `json:"password,string"`
}

type Work struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	user := User{
		Email:    "123@123",
		Password: 111,
	}

	word := Work{
		Name:    "szy",
		Address: "CHINA",
	}

	//b, err := json.Marshal(user)
	b, err := json.Marshal(struct {
		User
		Token    string `json:"token"`
		Password bool   `json:"password,omitempty"`
	}{
		User:  user,
		Token: "token",
	})
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Println(string(b))

	b, err = json.Marshal(struct {
		User
		Work
	}{user, word})
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Println(string(b))

	merge := []byte(`{
		"email": "111",
		"name": "222",
		"address": "333",
		"password": "13"
		}`)

	user1 := &User{}
	work1 := &Work{}
	err = json.Unmarshal(merge, &struct {
		*User
		*Work
	}{user1, work1})
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	fmt.Println(user1)
	fmt.Println(work1)
}
