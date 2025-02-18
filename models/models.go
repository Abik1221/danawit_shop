package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	First_Name     *string            `json:"first_tast" validate:"required"`
	Last_Name      *string            `json:"last_name" validate:"required"`
	Password       *string            `json:"password" validate:"password, required, min=8, max=12"`
	Email          *string            `json:"email" validate:"email"`
	Phone          *string            `json:"phone" validate:"required, phone"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `json:"refresh_token"`
	Created_at     time.Time          `json:"time"`
	Updated_at     time.Time          `json:"time"`
	User_id        *string            `json:"user_id"`
	User_Cart      []ProductUser      `jon:"user_cart"`
	Adress_Details []Address          `json:"adress"`
	Order_status   []Order            `json:"order_status"`
}

type Product struct {
	Product_ID    primitive.ObjectID `bson:"_id"`
	Product_name  *string            `json:"product_name"`
	Product_price *string            `json:"product_price"`
	Rating        int                `json:"rating"`
	Product_Image *string            `json:"product_image"`
}

type ProductUser struct {
	Product_ID    primitive.ObjectID `bson:"_id"`
	Product_name  *string            `json:"product_name"`
	Product_price *string            `json:"product_price"`
	Rating        *string            `json:"rating"`
	Product_Image *string            `json:"product_image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house"`
	Street     *string            `json:"street"`
	City       *string            `json:"city"`
	Pincode    *string            `json:"pincode"`
}

type Order struct {
	Order_id      primitive.ObjectID `bson:"_id"`
	Order_cart    *string            `json:"order_cart"`
	Ordered_at    time.Time          `json:"ordered_at"`
	Price         int                `jaon:"price"`
	Discount      int                `json:"discount"`
	PaymentMethod *string            `json:"paymentmethod"`
}

type Payment struct {
	Digital bool `json:"digital"`
	CDO     bool `json:"cdo"`
}
