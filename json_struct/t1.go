package main

import (
	"encoding/json"
	"fmt"
)

type Tmsg struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
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

type tuser struct {
	Tid      int    `json:"id"`
	Realname string `json:"realname"`
}

type responseDataf struct {
	Tmsg
	Tuser *tuser `json:"user"`
}

func main() {

	user := tuser{
		Tid:      1,
		Realname: "aaaa",
	}

	msg := Tmsg{
		ID:      2,
		Message: "sssss",
	}
	data := responseDataf{
		msg,
		&user,
	}
	message, _ := json.Marshal(data)

	fmt.Println(string(message))
}
