package model

import "time"

type Cart struct {
	ID         string     `json:"id" bson:"_id,omitempty"`
	UserID     string     `json:"userId" bson:"user_id"`
	Items      []CartItem `json:"items" bson:"items"`
	TotalPrice float64    `json:"totalPrice" bson:"total_price"`
	Status     string     `json:"status" bson:"status"`
	CreatedAt  time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time  `json:"updatedAt" bson:"updated_at"`
}