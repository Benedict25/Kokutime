package controllers

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Admin struct {
	Id_Admin int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id_User  int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type Drink struct {
	Id_Drink    int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type DrinksResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Drink `json:"data"`
}

type PromoCode struct {
	Id_PromoCode     int `json:"id"`
	Minimal_Purchase int `json:"minimal_purchase"`
	Promo_Amount     int `json:"promo_amount"`
	Quantity         int `json:"quantity"`
}

type PromoCodesResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []PromoCode `json:"data"`
}

type DetailedCarts struct {
	Id_Detailed_Cart int    `json:"id_detailed_cart"`
	Id_Cart          string `json:"id_cart"`
	Id_Drink         string `json:"id_drink"`
	Quantity         string `json:"quantity"`
}
