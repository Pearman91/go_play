package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	// '_ jmeno_balicku'  slouzi k importu balicku pro jeho side-effecty
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/gorilla/mux"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Email string
}

// vytvoreni databaze
func InitialMigration() {
//TODO you should just be opening your database once in your program, not every single time a method gets called. Once you've connected to the database in the InitialMigration method, you can use the same global db reference in other methods. Put the single defer close call in the same method as your HTTP router so that the database remains open as long as the web server is running.
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Nenavazano spojeni s databazi!")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Nenavazano spojeni s databazi!")
	}
	defer db.Close()
	
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Nenavazano spojeni s databazi!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name:name, Email: email})
	fmt.Fprintf(w, "New user added")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Nenavazano spojeni s databazi!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "User updated")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Nenavazano spojeni s databazi!")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User deleted")
}
