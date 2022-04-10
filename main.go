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

	// User Controller - User
	router.HandleFunc("/user-user", c.Authenticate(c.GetUserProfile, 1)).Methods("GET")
	router.HandleFunc("/user-user", c.Authenticate(c.EditProfile, 1)).Methods("PUT")

	// Drink Controller - Admin
	router.HandleFunc("/drink-admin", c.Authenticate(c.GetDrinks, 0)).Methods("GET")
	router.HandleFunc("/drink-admin", c.Authenticate(c.AddDrinks, 0)).Methods("POST")
	router.HandleFunc("/drink-admin", c.Authenticate(c.DeleteDrink, 0)).Methods("DELETE")
	router.HandleFunc("/drink-admin", c.Authenticate(c.UpdateDrink, 0)).Methods("PUT")

	// Promo Code Controller - Admin
	router.HandleFunc("/promo-code-admin", c.Authenticate(c.GetPromoCode, 0)).Methods("GET")
	router.HandleFunc("/promo-code-admin", c.Authenticate(c.AddPromoCode, 0)).Methods("POST")
	router.HandleFunc("/promo-code-admin", c.Authenticate(c.DeletePromoCode, 0)).Methods("DELETE")
	router.HandleFunc("/promo-code-admin", c.Authenticate(c.UpdatePromoCode, 0)).Methods("PUT")

	// Drink Controller - User
	router.HandleFunc("/drink-user", c.Authenticate(c.SeeDrinks, 1)).Methods("GET")
	router.HandleFunc("/drink-detail-user", c.Authenticate(c.SeeDetailedDrinks, 1)).Methods("GET")

	// Checkout Controller - User
	router.HandleFunc("/checkout-user", c.Authenticate(c.GetTotalPrice, 1)).Methods("GET")
	router.HandleFunc("/checkout-user", c.Authenticate(c.Checkout, 1)).Methods("POST")

	// Transaction - Admin
	router.HandleFunc("/transaction-admin", c.Authenticate(c.SalesReport, 0)).Methods("GET")
	router.HandleFunc("/transaction-admin", c.Authenticate(c.StatusManagement, 0)).Methods("PUT")
	router.HandleFunc("/transaction-user", c.Authenticate(c.SeeOrder, 1)).Methods("GET")
	router.HandleFunc("/detail-transaction-user", c.Authenticate(c.SeeDetailOrder, 1)).Methods("GET")
  
	// User Controller - Admin
	router.HandleFunc("/user-admin", c.Authenticate(c.SeeUsers, 0)).Methods("GET")
	router.HandleFunc("/user-admin", c.Authenticate(c.UpdateUser, 0)).Methods("PUT")
	router.HandleFunc("/user-admin", c.Authenticate(c.DeleteUser, 0)).Methods("DELETE")

	//ShoppingCartManagement
	router.HandleFunc("/cart-user", c.Authenticate(c.SeeCart, 1)).Methods("GET")
	router.HandleFunc("/cart-user", c.Authenticate(c.InsertCart, 1)).Methods("POST")
	router.HandleFunc("/cart-user", c.Authenticate(c.UpdateQuantity, 1)).Methods("PUT")
	router.HandleFunc("/cart-user", c.Authenticate(c.DeleteCart, 1)).Methods("DELETE")

	// Connection Notif
	http.Handle("/", router)
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
