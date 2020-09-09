package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quote struct
type Quote struct {
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	Customer   *Customer
	CustomerID primitive.ObjectID `bson:"customerID" json:"customerID"`
	Features   string             `json:"features,omitempty"`
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	ItemIds    *ItemIds           `bson:"items"`
	Items      *Items             `bson:"is"`
	JobsheetID primitive.ObjectID `bson:"jobsheetID" json:"jobsheetID"`
	Number     int                `bson:"number" json:"number"`
	Fees       struct {
		TotalCost   float64 `bson:"total"`
		Outstanding float64 `bson:"outstanding"`
	} `bson:"quotePrice"`
	Revision  int       `bson:"version" bson:"version"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// Address struct
type Address struct {
	Associate  string             `bson:"associate" json:"associate"`
	City       string             `bson:"city" json:"city"`
	CustomerID primitive.ObjectID `bson:"customerID" json:"customerID"`
	PostalCode string             `bson:"postalCode" json:"postalCode"`
	Province   string             `bson:"provinceCode" json:"province"`
	Street1    string             `bson:"street1" json:"street1"`
	Type       string             `bson:"type" json:"type"`
}

// Customer struct
type Customer struct {
	Email string `bson:"email" json:"email"`
	Name  struct {
		First  string `bson:"first" json:"first"`
		Last   string `bson:"last" json:"last"`
		Spouse string `bson:"spouse" json:"spouse,omitempty"`
	}
	Notes    string `bson:"notes" json:"notes"`
	Phones   []*Phone
	Address  *Address
	PhoneMap map[string]string
}

// Dim struct
type Dim struct {
	Decimal  float64 `bson:"decimal" json:"decimal"`
	Fraction string  `bson:"fraction,omitempty" json:"fraction"`
	Inch     int     `bson:"inch" json:"inch"`
}

// Dims struct
type Dims struct {
	Height *Dim `bson:"height" json:"height"`
	Width  *Dim `bson:"width" json:"width"`
}

// Group struct
type Group struct {
	Dims  *Dims          `bson:"dims" json:"dims"`
	Items []*GroupWindow `bson:"items" json:"items"`
	Specs bson.M         `bson:"specs" json:"specs"`
	Qty   int            `bson:"qty" json:"qty"`
	Rooms []string       `bson:"rooms" json:"rooms"`
}

// GroupWindow struct
type GroupWindow struct {
	Dims    *Dims  `bson:"dims" json:"dims"`
	Product bson.M `bson:"product" json:"product"`
	Qty     int    `bson:"qty" json:"qty"`
	Specs   bson.M `bson:"specs" json:"specs"`
}

// ItemIds struct
type ItemIds struct {
	Group  []string `bson:"group" bson:"group"`
	Other  []string `bson:"other" bson:"other"`
	Window []string `bson:"window" bson:"window"`
}

// Items struct
type Items struct {
	Group  []*Group
	Other  []*Other
	Window []*Window
}

// JobSheet struct
type JobSheet struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Features string             `bson:"features" json:"features,omitempty"`
}

// Other struct
type Other struct {
	Description string   `bson:"description" json:"description"`
	Qty         int      `bson:"qty" json:"qty"`
	Rooms       []string `bson:"rooms" json:"rooms"`
	Specs       struct {
		Options  string `bson:"options" json:"options"`
		Location string `bson:"location" json:"location,omitempty"`
	}
}

// Phone struct
type Phone struct {
	CountryCode string `bson:"countryCode" json:"countryCode"`
	Number      string `bson:"number" json:"number"`
	Type        string `bson:"_id" json:"type"`
}

// Product struct
type Product struct {
	Name string `bson:"name" json:"name"`
}

// Window struct
type Window struct {
	Dims        *Dims              `bson:"dims" json:"dims"`
	Qty         int                `bson:"qty" json:"qty"`
	ProductID   primitive.ObjectID `bson:"productID" json:"productID"`
	ProductName string
	Rooms       []string `bson:"rooms" json:"rooms"`
	Specs       bson.M   `bson:"specs" json:"specs"`
}
