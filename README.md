# Delivery API for Golang

## Summary

The Delivery Golang API is used for retrieving content from MariaDB. The API is currently using Golang version 1.14, Go SQL driver 1.5.0 and MariaDB 10.4.17. 
To check the response I used Postman.

## Installation

1) Run main.go. 
2) Import `orders.sql` file in the mySQL to create the database and tables

## Improvements

1) Location table can be created in which we can get data from google MAPs API to show the current location of the orders.
3) Some error cases and exception handling to be done.
4) Return of orders can be implemented,so status of that can be added.

## Endpoints for Delivery_API

The endpoints work on `localhost:8080`.

### Create an Order

Use `POST` with endpoint `/orders`. Using this, we can create an order.

## Payload for Creating an Order
  ```
  {
    "customerName": "Harry",
    "orderedAt": "2019-11-09T21:21:46Z",
    "items": [
        {
            "itemCode": "123",
            "description": "IPhone 11X",
            "quantity": 1
        }
    ],
    "delivery":[
        {
            "location":3,
            "status_id":1
        }
    ]
}
```

### Retreive an Order

Use `GET` with endpoint `/orders/{orderId}`. Using this, we can retreive an order with the Order ID.

### Retreive all Orders

Use `GET` with endpoint `/orders/all`. Using this, we can retreive all orders.

## Endpoints for Order_Status

The endpoints work on `localhost:8000`.

Use `GET` with endpoint `/orders/delivery/status/{orderId}`. Using this, we can fetch the order status with the Order ID.

## Endpoints for Update_Status

The endpoints work on `localhost:9000`.

Use `PUT` with endpoint `/orders/{deliveryId}`. Using this, we can update the order delivery status with the delivery ID.

## Status_id Code
1) By default `status_id = 1` is send while creating a order it denotes `Processing` of order.
2) `status_id = 2` is for order is in `Pickup` status.
3) `status_id =3` is for order is already `Delivered` status.

>order status is changed by delivery partner.

## Payload for Update_Status
```
        {
            "deliveryId":1,
            "location":1,
            "order_id":1,
            "status_id":3
        }
```

## Postman Collection

To import the collection in Postman, use this link:

`https://www.postman.com/collections/0602b6b48f6edbb1cbb3`
