package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Domain model
*/

type Jar struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Category    string             `bson:"category" json:"category"`
	Price       float64            `bson:"price" json:"price"`
	StockQty    int                `bson:"stock_qty" json:"stock_qty"`
	ImageUrl    string             `bson:"image_url" json:"image_url"`
	Attributes  JarAttributes      `bson:"attributes" json:"attributes"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type JarAttributes struct {
	ClayType       string `bson:"clay_type" json:"clay_type"`
	Dimensions     string `bson:"dimensions" json:"dimensions"`
	Capacity       string `bson:"capacity" json:"capacity"`
	Weight         string `bson:"weight" json:"weight"`
	FoodSafe       bool   `bson:"food_safe" json:"food_safe"`
	MicrowaveSafe  bool   `bson:"microwave_safe" json:"microwave_safe"`
	DishwasherSafe bool   `bson:"dishwasher_safe" json:"dishwasher_safe"`
	GlazeType      string `bson:"glaze_type,omitempty" json:"glaze_type,omitempty"`
	ProductionType string `bson:"production_type" json:"production_type"`
}

/*
DTOs Request model
*/

type CreateJarRequest struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Price       float64       `json:"price"`
	StockQty    int           `json:"stock_qty"`
	ImageURL    string        `json:"image_url"`
	Attributes  JarAttributes `json:"attributes"`
}

/*
Validation
*/

func (j *Jar) Validate() error {
	switch {
	case j.Name == "":
		return errors.New("Name attribute is mandatory")
	case len(j.Name) > 200:
		return errors.New("Name attribute must be less than 200 characters")
	case j.Price < 0.01:
		return errors.New("Price must be at least 0.01")
	case j.Price > 10000:
		return errors.New("Price must not exceed 10000")
	case j.StockQty < 0:
		return errors.New("Stock quantity cannot be negative")
	case j.StockQty > 100000:
		return errors.New("Stock quantity exceeds allowed maximum")
	}
	return nil
}

/*
Lifecycle Hooks
*/

func (j *Jar) PrepareForCreate() {
	if j.ID.IsZero() {
		j.ID = primitive.NewObjectID()
	}
	now := time.Now().UTC()
	j.CreatedAt = now
	j.UpdatedAt = now
}

func (j *Jar) PrepareForUpdate() {
	j.UpdatedAt = time.Now().UTC()
}

/*
Domain Events
*/

type JarEvent struct {
	Type      string    `json:"type"`
	JarID     string    `json:"jar_id"`
	Payload   *Jar      `json:"payload,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
