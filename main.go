package main

import (
	c "Kokutime/controllers"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Connection
	"github.com/gorilla/mux"           // Router
)

func main() {

	// 0 = Admin ; 1 = User

	// Processing, Delivering, Received

	router := mux.NewRouter()

	// End Point

	// Register
	router.HandleFunc("/register-user", c.UserRegister).Methods("POST")

	// Login
	router.HandleFunc("/login-user", c.UserLogin).Methods("POST")
	router.HandleFunc("/login-admin", c.AdminLogin).Methods("POST")

	// Logout
	router.HandleFunc("/logout", c.Logout).Methods("POST")

	// User Controller
	router.HandleFunc("/user-admin", c.Authenticate(c.GetUsers, 0)).Methods("GET")
	router.HandleFunc("/user-user", c.Authenticate(c.GetUserProfile, 1)).Methods("GET")
	router.HandleFunc("/user-user", c.Authenticate(c.EditProfile, 1)).Methods("PUT")

	// Drink's Controller - Admin
	router.HandleFunc("/drink-admin", c.Authenticate(c.GetDrinks, 0)).Methods("GET")
	router.HandleFunc("/drink-admin", c.Authenticate(c.AddDrinks, 0)).Methods("POST")
	router.HandleFunc("/drink-admin", c.Authenticate(c.DeleteDrink, 0)).Methods("DELETE")
	router.HandleFunc("/drink-admin", c.Authenticate(c.UpdateDrink, 0)).Methods("PUT")

	// Connection Notif
	http.Handle("/", router)
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
