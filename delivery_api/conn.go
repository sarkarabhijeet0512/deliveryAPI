package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	// On the log mod for sql queries in console
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("CREATE DATABASE orders")
	db.Exec("USE orders")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&Order{}, &Item{}, &Delivery{}, &Status{})
}
