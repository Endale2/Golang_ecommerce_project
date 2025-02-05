package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	First_Name    *string            `json:"first_name" bson:"first_name,omitempty"`
	Last_Name     *string            `json:"last_name" bson:"last_name,omitempty"`
	Password      *string            `json:"password,omitempty" bson:"password,omitempty"`
	Email         *string            `json:"email" bson:"email,omitempty"`
	Phone         *string            `json:"phone" bson:"phone,omitempty"`
	Token         *string            `json:"token,omitempty" bson:"token,omitempty"`
	Refresh_Token *string            `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
	Created_At    time.Time          `json:"created_at" bson:"created_at"`
	Updated_At    time.Time          `json:"updated_at" bson:"updated_at"`
	User_ID       *string            `json:"user_id" bson:"user_id,omitempty"`
	User_Cart     []ProductUser      `json:"user_cart" bson:"user_cart"`
	Address       []Address          `json:"address" bson:"address"`
	Order_Status  []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_Id   primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	Product_Name *string            `json:"product_name" bson:"product_name,omitempty"`
	Price        *uint64            `json:"price" bson:"price,omitempty"`
	Rating       *uint8             `json:"rating" bson:"rating,omitempty"`
	Image        *string            `json:"image" bson:"image,omitempty"`
}

type ProductUser struct {
	Product_Id   primitive.ObjectID `json:"product_id" bson:"product_id,omitempty"`
	Product_Name *string            `json:"product_name" bson:"product_name,omitempty"`
	Price        *uint64            `json:"price" bson:"price,omitempty"`
	Rating       *uint8             `json:"rating" bson:"rating,omitempty"`
	Image        *string            `json:"image" bson:"image,omitempty"`
}

type Address struct {
	Address_Id primitive.ObjectID `json:"address_id" bson:"address_id,omitempty"`
	Home       *string            `json:"home" bson:"home,omitempty"`
	Street     *string            `json:"street" bson:"street,omitempty"`
	City       *string            `json:"city" bson:"city,omitempty"`
	Pincode    *string            `json:"pincode" bson:"pincode,omitempty"`
}

type Order struct {
	Order_Id       primitive.ObjectID `json:"order_id" bson:"order_id,omitempty"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          *uint64            `json:"price" bson:"price,omitempty"`
	Discount       *uint64            `json:"discount" bson:"discount,omitempty"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"` // Cash on Delivery
}
