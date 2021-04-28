package main

import "time"

type Order struct {
	// gorm.Model
	OrderID      uint       `json:"orderId" gorm:"primary_key"`
	CustomerName string     `json:"customerName"`
	OrderedAt    time.Time  `json:"orderedAt"`
	Items        []Item     `json:"items" gorm:"foreignkey:OrderID"`
	Delivery     []Delivery `json:"delivery" gorm:"foreignkey:OrderID"`
}
type Item struct {
	// gorm.Model
	LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
type Delivery struct {
	DeliveryID uint `json:"deliveryId" gorm:"primary_key"`
	Location   uint `json:"location"`
	StatusID   uint `json:"status_id" gorm:"foreignkey:StatusID"`
	OrderID    uint `json:"-"`
}
type Status struct {
	StatusID  uint      `json:"statusId" gorm:"primary_key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdat"`
}
type Demo struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	LineItemID   uint      `json:"lineItemId" gorm:"primary_key"`
	ItemCode     string    `json:"itemCode"`
	Description  string    `json:"description"`
	Quantity     uint      `json:"quantity"`
	DeliveryID   uint      `json:"deliveryId" gorm:"primary_key"`
	Location     uint      `json:"location"`
	StatusID     uint      `json:"status_id" gorm:"foreignkey:StatusID"`
	OrderID      uint      `json:"-"`
	Value        string    `json:"value"`
	CreatedAt    time.Time `json:"createdat"`
}
