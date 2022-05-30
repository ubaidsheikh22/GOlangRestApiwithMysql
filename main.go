package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initializeRouter() {
	rout := mux.NewRouter()

	rout.HandleFunc("/getallusers", GetAllUsers).Methods("GET")
	rout.HandleFunc("/getuser/{id}", GetSingleuser).Methods("GET")
	rout.HandleFunc("/createuser", CreateUser).Methods("POST")
	rout.HandleFunc("/updateuser/{id}", Updateuser).Methods("PUT")
	rout.HandleFunc("/deleteuser/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", rout))
}

var DB *gorm.DB
var err error

const DSN = "root:root@tcp(127.0.0.1:3306)/UserDetails?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Connecting Database Error", err.Error())
	}
	DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	UserName string `json:"UserName"`
	Email    string `json:"Email"`
	Password string `Json:"Password"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetSingleuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("conetent-type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func Updateuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("conetent-type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("User Deleted Successfully")
}

func main() {
	InitialMigration()
	initializeRouter()

}
