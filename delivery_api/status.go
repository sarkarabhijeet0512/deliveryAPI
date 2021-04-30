package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var status Demo
	db.Table("orders").
		Select("orders.customer_name,orders.mobile_number,deliveries.order_id,deliveries.location,statuses.value ").
		Joins("JOIN items on items.order_id=orders.order_id").
		Joins("JOIN deliveries on deliveries.order_id=orders.order_id").
		Joins("JOIN statuses on statuses.status_id=deliveries.status_id").
		Where("orders.order_id=?", params["orderId"]).Find(&status)
	res := make(map[string]interface{})
	res["customer_name"] = status.CustomerName
	res["contact_no"] = status.MobileNumber
	res["delivery_status"] = status.Value
	res["location"] = status.Location

	des, _ := json.Marshal(res)
	w.Write([]byte(des))
}
