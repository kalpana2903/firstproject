package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getuser(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "kalpana",
		Price: 10.0,
	}

	json.NewEncoder(w).Encode(user)
}
func adduser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in request")
	id := r.FormValue("id")
	name := r.FormValue("name")
	price := r.FormValue("price")
	fmt.Println(id, name, price)
}

func main() {
	http.HandleFunc("/adduser", adduser)
	http.HandleFunc("/getuser", getuser)     // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
