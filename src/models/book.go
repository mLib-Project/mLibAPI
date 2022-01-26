package models

import primitive "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id          primitive.ObjectID `form:"_id" json:"_id" validate:"required"`
	Family      string             `form:"family" json:"family,omitempty" validate:"required"`
	Category    string             `form:"category" json:"category,omitempty" validate:"required"`
	Subcategory string             `form:"subcategory" json:"subcategory,omitempty" validate:"required"`
	Name        string             `form:"name" json:"name,omitempty" validate:"required"`
	Author      string             `form:"author" json:"author,omitempty" validate:"required"`
	Source      string             `form:"source" json:"source,omitempty" validate:"required"`
	// ID          string             `form:"ID" json:"ID"`
}
