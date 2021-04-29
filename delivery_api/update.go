package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func updateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder DeliveryUpdate
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	delivery_status := updatedOrder.StatusID
	Otp_status := updatedOrder.Otp

	if Otp_status != "" && delivery_status == 3 {
		db.Model(&updatedOrder).Where("delivery_id = ?", params["deliveryId"]).Update(&updatedOrder)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedOrder)
	} else if delivery_status == 1 || delivery_status == 2 {
		db.Model(&updatedOrder).Where("delivery_id = ?", params["deliveryId"]).Update(&updatedOrder)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedOrder)
	} else {
		w.Write([]byte(`{"response":"please provide otp"}`))

	}
}
