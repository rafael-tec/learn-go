package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Account struct {
	Number int `json:"number_pk"`
	ID string `json:"id_pk"`
}

func main() {
	account := Account{Number: 9482, ID: "a3d5g8n1ad9"}
	otherAccount := Account{Number: 3290, ID: "x0a4h9s4g73"}

	marshalToJson(account)
	marshalToJsonEncoded(otherAccount)
	unmarshalToJson([]byte(`{"number_pk": 2931, "id_pk": "g1s4x3m9zg0"}`))
}

func unmarshalToJson(j []byte) {
	var account Account
	
	err := json.Unmarshal(j, &account)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(account)
}

func marshalToJson(a any) {
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}

func marshalToJsonEncoded(a any) {
	err := json.NewEncoder(os.Stdout).Encode(a)
	if err != nil {
		fmt.Println(err)
	}
}