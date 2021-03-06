package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedOrder Delivery
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	delivery_status := updatedOrder.StatusID
	Otp_status := updatedOrder.Otp
	err := db.Raw("SELECT otp,delivery_id FROM deliveries where delivery_id = ?", params["deliveryId"]).
		Find(&updatedOrder).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			w.Write([]byte(`{"response":"` + "no data found" + `"}`))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"response":"` + "oops! something went wrong! Try again" + `"}`))
		}
	} else {
		if updatedOrder.Otp == Otp_status {
			if Otp_status != 0 && delivery_status == 3 {
				db.Model(&updatedOrder).
					Where("delivery_id = ?", params["deliveryId"]).Update(&updatedOrder)
				res := make(map[string]interface{})
				res["response"] = "otp verification success"
				res["status"] = "Order Delivered"
				des, _ := json.Marshal(res)
				w.Write([]byte(des))
			}
		} else if delivery_status == 1 || delivery_status == 2 {
			db.Model(&updatedOrder).Where("delivery_id = ?", params["deliveryId"]).Update(&updatedOrder)

			res := make(map[string]interface{})
			res["response"] = "order status updated successfully"
			if delivery_status == 1 {
				res["status"] = "processing"
			} else {
				res["status"] = "Picked up"
			}
			des, _ := json.Marshal(res)
			w.Write([]byte(des))
		} else {
			w.Write([]byte(`{"response":"please provide valid otp"}`))

		}
	}
}
