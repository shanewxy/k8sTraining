package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("sqlite3", "./training.db")
	checkError(err)
	defer db.Close()
	err = db.Ping()
	db.Exec("CREATE table if not exists User (id INTEGER PRIMARY KEY NOT NULL,username text,password text)")
	checkError(err)
	port := strconv.Itoa(9999)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.ListenAndServe(":"+port, nil)
	fmt.Println("Service up")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func createHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.DefaultCost)
	checkError(err)
	_, err = db.Exec(`Insert into User (username,password) values(?,?)`, username, hashedPassword)
	checkError(err)
	fmt.Println("Created user: ", username)
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * from user ")
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		fmt.Println("User: ", user)
	}
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	checkError(err)
	_, err = db.Exec(`update user set password=? where username=?`, hashedPassword, username)
	fmt.Println("Updated password for ", username)
	checkError(err)
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	checkError(err)
	_, err = db.Exec(`delete from user where id=?`, id)
	checkError(err)
	fmt.Println("delete user whose id = ", id)
}
