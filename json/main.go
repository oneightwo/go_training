package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UsersModel struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Age    int     `json:"age"`
	Social *Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func (u *UsersModel) print() {
	fmt.Println("users:")
	for i := range u.Users {
		fmt.Println("	name:", u.Users[i].Name)
		fmt.Println("	type:", u.Users[i].Type)
		fmt.Println("	age:", u.Users[i].Age)
		fmt.Println("	social:")
		fmt.Println("		facebook:", u.Users[i].Social.Facebook)
		fmt.Println("		twitter:", u.Users[i].Social.Twitter)
	}
}

func main() {

	file, errorFile := ioutil.ReadFile("example.json")

	if errorFile != nil {
		fmt.Println(errorFile)
	} else {
		var u = UsersModel{}
		errorJson := json.Unmarshal([]byte(file), &u)

		if errorJson != nil {
			fmt.Println(errorJson)
		} else {
			u.print()
		}
	}
}
