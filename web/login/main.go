package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type server struct {
	db *sql.DB
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type UserOutput struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Login      string `json:"login"`
	Error      string `json:"error"`
}

type IncomingData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Error struct {
	Message string `json:"error"`
}

type Registered struct {
	Registered int
}

func (s *server) join(w http.ResponseWriter, r *http.Request) {
	user := User{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), &user)
	checkError(err)
	login := user.Login
	if isRegistered(s, login) {
		_, _ = w.Write(getUser(s, login, "the user is already registered"))
	} else {
		bytesPassword := []byte(user.Password)
		hashPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
		checkError(err)
		user.Password = string(hashPassword)
		insertUser(s, user)
		_, _ = w.Write(getUser(s, login, "null"))
	}
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	incomingData := IncomingData{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal([]byte(body), &incomingData)
	checkError(err)
	if isRegistered(s, incomingData.Login) {
		if passwordVerification(s, incomingData.Login, incomingData.Password) {
			_, _ = w.Write(getUser(s, incomingData.Login, "null"))
		} else {
			_, _ = w.Write(getError("invalid password"))
		}
	} else {
		_, _ = w.Write(getError("invalid login"))
	}
}

func passwordVerification(s *server, login string, password string) bool {
	user := User{}
	rows := s.db.QueryRow("SELECT * FROM user_account WHERE login = $1", login)
	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Age, &user.Login, &user.Password)
	bytesPassword := []byte(password)
	bytesHashPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(bytesHashPassword, bytesPassword)
	if err == nil {
		return true
	} else {
		return false
	}
}

func getError(textError string) []byte {
	var error Error
	error.Message = textError
	bytes, err := json.Marshal(error)
	checkError(err)
	return bytes
}

func insertUser(s *server, user User) {
	_, err := s.db.Exec("INSERT INTO user_account (firstname, lastname, age, login, password) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Age, user.Login, user.Password)
	checkError(err)
}

func getUser(s *server, login string, error string) []byte {
	userOutput := UserOutput{}
	rows := s.db.QueryRow("SELECT * FROM user_account WHERE login = $1", login)
	err := rows.Scan(&userOutput.Id, &userOutput.FirstName, &userOutput.LastName, &userOutput.Age, &userOutput.ExternalId, &userOutput.Login, &userOutput.Error)
	userOutput.Error = error
	bytes, err := json.Marshal(userOutput)
	checkError(err)
	return bytes
}

func isRegistered(s *server, login string) bool {
	row := s.db.QueryRow("SELECT COUNT(*) FROM user_account WHERE login =  $1", login)
	registered := Registered{}
	err := row.Scan(&registered.Registered)
	checkError(err)
	if registered.Registered == 0 {
		return false
	} else {
		return true
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dataSource := "user=postgres password=lumia640 dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	server := server{db: db}
	checkError(err)
	router := mux.NewRouter()
	router.HandleFunc("/join", server.join).Methods("POST")
	router.HandleFunc("/login", server.login).Methods("POST")
	http.Handle("/", router)
	err = http.ListenAndServe(":1024", nil)
	checkError(err)
}
