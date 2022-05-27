package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Observations struct {
	IsFeatured        string
	IsFixed           string
	IsMessage         string
	Message           string
	IsBussinesGravanz string
	IsPort            string
}

type PlaceOfDelivery struct {
	Zone string
	Port string
}

type Product struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductName     string             `bson:"produc_name,omitempty"`
	Status          string             `bson:"status,omitempty"`
	BusinessType    string             `bson:"business_type,omitempty" json:"business_type,omitempty"`
	DeliveryPeriod  string             `bson:"delivey_period,omitempty" json:"delivey_period,omitempty"`
	Quality         string             `bson:"quality,omitempty" json:"quality,omitempty"`
	Price           string             `bson:"price,omitempty" json:"price,omitempty"`
	TypeCoin        string             `bson:"type_coin,omitempty" json:"type_coin,omitempty"`
	WayPay          string             `bson:"way_pay,omitempty" json:"way_pay,omitempty"`
	Expiration      string             `bson:"expiration,omitempty" json:"expiration,omitempty"`
	Bonus           string             `bson:"bonus,omitempty" json:"bonus,omitempty"`
	Reduction       string             `bson:"reduction,omitempty" json:"reduction,omitempty"`
	Pesification    string             `bson:"pesification,omitempty" json:"pesification,omitempty"`
	HourP           string             `bson:"hour_p,omitempty" json:"hour_p,omitempty"`
	PriceId         string             `bson:"price_id,omitempty" json:"price_id,omitempty"`
	Harvest         string             `bson:"harvest,omitempty" json:"harvest,omitempty"`
	Observations    Observations       `bson:"inline"`
	PlaceOfDelivery PlaceOfDelivery    `bson:"inline"`
	CreateAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Product           Product            `bson:"inline"`
	SellerName        string             `bson:"seller_name,omitempty" json:"seller_name,omitempty"`
	StatusOrder       string             `bson:"status_order,omitempty" json:"status_order,omitempty"`
	HourO             string             `bson:"hour_o,omitempty" json:"hour_o,omitempty"`
	Particularidades  string             `bson:"particularidades,omitempty" json:"particularidades,omitempty"`
	BusinessSpecifics string             `bson:"business_specifics,omitempty" json:"business_specifics,omitempty"`
	Tons              string             `bson:"tons,omitempty" json:"tons,omitempty"`
}
