package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Delivery struct {
	DeliveryID uint `json:"deliveryId" gorm:"primary_key"`
	Location   uint `json:"location"`
	OrderID    uint `json:"order_id"`
	StatusID   uint `json:"status_id"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("USE orders")

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders/{deliveryId}", updateOrder).Methods("PUT")
	initDB()
	log.Fatal(http.ListenAndServe(":9000", router))
}
func updateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder Delivery
	params := mux.Vars(r)
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	db.Model(&updatedOrder).Where("delivery_id = ?", params["deliveryId"]).Update(&updatedOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)
}
