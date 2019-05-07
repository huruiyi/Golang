package main

import (
	"./lib"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var database *sql.DB

type Users struct {
	Users []lib.User `json:"users"`
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	NewUser := lib.User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:8080/api/users?start="+string(next)+"; rel=\"next\"")

	rows, _ := database.Query("select * from users LIMIT 10")
	Response := Users{}

	for rows.Next() {

		user := lib.User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)

		Response.Users = append(Response.Users, user)
	}

	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

func main() {
	db, err := sql.Open("mysql", "root@/social_network")
	if err != nil {

	}
	database = db
	routes := mux.NewRouter()
	routes.HandleFunc("/api/users", UserCreate).Methods("POST")
	routes.HandleFunc("/api/users", UsersRetrieve).Methods("GET")
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)
}
