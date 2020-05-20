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
	Tg *Tmsg `json:"Tmsg"`
	Ter *Tuser `json:"Tuser"`
}

func (me Tmsg) MarshalJSON() ([]byte, error) {
	type Alias Tmsg
	return json.Marshal(&struct {
		Message string `json:"message"`
		Alias
	}{
		Message: "1",
		Alias:   (Alias)(me),
	})
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
		Tg:&msg,
		Ter:&user,
	}
	message, _ := json.Marshal(data)
	fmt.Println(string(message))
}