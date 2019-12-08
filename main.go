package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
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
	//建表
	db.Exec("CREATE table if not exists User (id INTEGER PRIMARY KEY NOT NULL,username text,password text)")
	checkError(err)
	port := strconv.Itoa(9999)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/", indexHandler)
	fmt.Println("Service up")

	http.ListenAndServe(":"+port, nil)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")
	data := map[string]string{"name": "wxy"}
	t.Execute(w, data)
}

//创建用户
func createHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.DefaultCost)
	checkError(err)
	_, err = db.Exec(`Insert into User (username,password) values(?,?)`, username, password)
	checkError(err)
	fmt.Printf("Created user: %s\n", username)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

//列出所有用户
func listHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, _ := db.Query("SELECT * from user ")
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	juser, _ := json.Marshal(users)
	fmt.Fprintf(w, string(juser))
}

//通过username，更新用户密码
func updateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	checkError(err)
	_, err = db.Exec(`update user set password=? where username=?`, password, username)
	fmt.Printf("Updated password for %s\n", username)
	checkError(err)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

//通过用户id删除用户
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	checkError(err)
	stmt := `delete from user where id=?`
	_, err = db.Exec(stmt, id)
	fmt.Println(stmt)
	checkError(err)
	fmt.Printf("delete user whose id = %d\n", id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
