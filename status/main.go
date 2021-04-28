package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Demo struct {
	Location  uint      `json:"location"`
	OrderID   uint      `json:"order_id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdat"`
}

var db *sql.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("USE orders")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders/{orderId}", getPost).Methods("GET")

	initDB()

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT deliveries.order_id,deliveries.location,statuses.value FROM deliveries INNER JOIN statuses ON deliveries.status_id=statuses.status_id where deliveries.order_id=?", params["orderId"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post Demo
	for result.Next() {
		err := result.Scan(&post.OrderID, &post.Location, &post.Value)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}
