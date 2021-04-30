package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	fmt.Println("starting the server.....")
	fmt.Println("use Port :8080")
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", CreateOrder).Methods("POST")
	// Read
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	// Update
	router.HandleFunc("/orders/{deliveryId}", updateOrder).Methods("PUT")
	// Getall Orders
	router.HandleFunc("/orders/all", getOrders).Methods("GET")
	//getdelivery status for delivery partners
	router.HandleFunc("/orders/delivery/status/{orderId}", getPost).Methods("GET")
	InitDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var results Demo
	//do error handling
	db.Table("orders").Select("*").
		Joins("JOIN items on items.order_id=orders.order_id").
		Joins("JOIN deliveries on deliveries.order_id=orders.order_id").
		Joins("JOIN statuses on statuses.status_id=deliveries.status_id").
		Where("orders.order_id=?", params["orderId"]).Find(&results)
	json.NewEncoder(w).Encode(results)
}
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []Demo
	db.Table("orders").Select("*").
		Joins("JOIN items on items.order_id=orders.order_id").
		Joins("JOIN deliveries on deliveries.order_id=orders.order_id").
		Joins("JOIN statuses on statuses.status_id=deliveries.status_id").
		Find(&results)
	json.NewEncoder(w).Encode(results)
}
