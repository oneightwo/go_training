package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

func (u *UsersModel) HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	var str string
	if len(r.URL.RawQuery) > 0 {
		str = r.URL.Query().Get("number")
		if str == "" {
			w.WriteHeader(400)
			return
		} else {
			number, _ := strconv.Atoi(str)
			if len(u.Users) >= number && number > 0 {
				number--
				text := u.serialize(number)
				_, _ = w.Write(text)
				w.WriteHeader(200)
			} else {
				w.WriteHeader(404)
				return
			}
		}
	} else {
		w.WriteHeader(404)
	}
}

func (u *UsersModel) serialize(number int) []byte {
	jsons, err := json.Marshal(u.Users[number])
	if err != nil {
		fmt.Println(err)
	}
	return jsons
}

func (u *UsersModel) deserialize() {
	file, errorFile := ioutil.ReadFile("example.json")
	if errorFile != nil {
		fmt.Println(errorFile)
	} else {
		errorJson := json.Unmarshal([]byte(file), &u)
		if errorJson != nil {
			fmt.Println(errorJson)
		}
	}
}

func main() {
	var u = UsersModel{}
	u.deserialize()
	http.HandleFunc("/users", u.HomeRouterHandler)
	err := http.ListenAndServe(":1024", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
