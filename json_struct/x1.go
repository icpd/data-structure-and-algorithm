package main

import (
	"encoding/json"
	"fmt"
)

type Tmsg struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}


type Tuser struct {
	ID      int    `json:"id"`
	Realname string `json:"realname"`
}

type responseDataf struct {
	Tg Tmsg `json:"Tmsg"`
	Ter Tuser `json:"Tuser"`
}

func main() {

	user := Tuser{
		ID:1,
		Realname: "aaaa",
	}

	msg := Tmsg{
		ID:2,
		Message: "sssss",
	}
	data := responseDataf{
		msg,
		user,
	}
	message, _ := json.Marshal(data)
	fmt.Println(string(message))
}