package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	var delivery Delivery
	json.NewDecoder(r.Body).Decode(&order)
	// some messaging syatem can be implemented to send to otp to the customers numbers
	rand := rand.Intn(10000)
	db.Create(&order)
	db.Model(&delivery).Where("order_id = ?", order.OrderID).Update("otp", rand)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
