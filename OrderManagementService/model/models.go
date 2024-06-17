package model

import "time"

type Order struct {
	OrderID         string
	UserID          string
	Timestamp       time.Time
	ItemID          string
	RiderID         string
	RestaurantID    string
	DeliveryAddress string
	Assigned        bool
	Delivered       bool
}

type Rider struct {
	ID       string
	Name     string
	Assigned bool
}
