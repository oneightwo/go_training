package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	ProductType string `json:"type"`
}

type server struct {
	db *sql.DB
}

func (s *server) productsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := s.db.Query("SELECT * FROM product")
	checkError(err)
	defer rows.Close()

	var products []Product

	for rows.Next() {
		product := Product{}

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.ProductType)

		checkError(err)

		products = append(products, product)
	}

	bytes, err := json.Marshal(products)
	checkError(err)
	_, _ = w.Write(bytes)
}

func (s *server) productByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product := Product{}

	rows, err := s.db.Query("SELECT * FROM product WHERE id = $1", id)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.ProductType)

		checkError(err)
	}

	bytes, err := json.Marshal(product)

	checkError(err)
	_, _ = w.Write(bytes)
}

func (s *server) productUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product := Product{}

	body, err := ioutil.ReadAll(r.Body)

	checkError(err)
	err = json.Unmarshal([]byte(body), &product)
	checkError(err)

	rows, err := s.db.Exec("UPDATE product SET name = $1, price = $2, type = $3 WHERE id = $4",
		product.Name, product.Price, product.ProductType, id)

	checkError(err)
	rows.RowsAffected()
}

func (s *server) productDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	rowsPurchase, errPurchase := s.db.Query("DELETE FROM purchase WHERE product_id = $1", id)
	rows, err := s.db.Exec("DELETE FROM product WHERE id = $1", id)

	checkError(err)
	checkError(errPurchase)

	rows.RowsAffected()
	defer rowsPurchase.Close()
}

func (s *server) productAddHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{}

	body, err := ioutil.ReadAll(r.Body)

	checkError(err)
	err = json.Unmarshal([]byte(body), &product)
	checkError(err)
	fmt.Print(string(body))
	rows, err := s.db.Exec("INSERT INTO product (name, price, type) VALUES ($1, $2, $3)",
		product.Name, product.Price, product.ProductType)

	checkError(err)
	rows.RowsAffected()
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
	router.HandleFunc("/products", server.productsHandler).Methods("GET")
	router.HandleFunc("/products/id/{id}", server.productByIdHandler).Methods("GET")
	router.HandleFunc("/products/id/{id}", server.productUpdateHandler).Methods("PUT")
	router.HandleFunc("/products/id/{id}", server.productDeleteHandler).Methods("DELETE")
	router.HandleFunc("/products/add", server.productAddHandler).Methods("POST")
	http.Handle("/", router)
	err = http.ListenAndServe(":1024", nil)
	checkError(err)
}
