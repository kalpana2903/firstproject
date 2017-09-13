package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	users    = "kalpanaverma"
	password = "kalpu@2903"
	dbname   = "super_awesome_application"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// type User1 struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

var db *sqlx.DB

func initializeDB() {

	d, err := sqlx.Open("postgres", "user=kalpanaverma password=kalpu@2903 host=localhost port=5432 dbname=super_awesome_application sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = d.Ping()
	if err != nil {
		panic(err)
	}

	db = d
	fmt.Println("initialize db")

}

func init() {
	initializeDB()
}

func adduser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("adduser")
	b, err := ioutil.ReadAll(r.Body)

	// Unmarshal
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		fmt.Print("Error")
	}

	log.Println(err, "======", u)

	sqlStatement := `
	INSERT INTO userinfo (ID,name)
	VALUES ($1, $2)
	RETURNING id`
	Id := 0
	err2 := db.QueryRow(sqlStatement, u.ID, u.Name).Scan(&Id)
	if err2 != nil {
		fmt.Print(err2.Error())
	}
	fmt.Println("New record ID is:", Id)

}

func getusers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Calling getusers")
	users := []User{}
	e := db.Select(&users, "SELECT * FROM userinfo")
	//_, e := db.Query("SELECT * FROM userinfo")
	if e != nil {
		fmt.Println(e.Error())
	}
	log.Println(e, "======", users)
	json.NewEncoder(w).Encode(users)

}

func getuser(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("id")
	fmt.Println("Calling getuser")

	var user User
	fmt.Println("Still connected")
	if ID != "" {
		fmt.Println(ID)
		e := db.Get(&user, `SELECT * FROM userinfo WHERE id=$1`, ID)
		if e != nil {
			fmt.Println(e.Error())
		}
		log.Println(e, "======", user)
		json.NewEncoder(w).Encode(user)

	}

}
func main() {

	http.HandleFunc("/adduser", adduser)
	http.HandleFunc("/getusers", getusers)
	http.HandleFunc("/getuser", getuser)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
