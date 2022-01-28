package main

import "gorm.io/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model

	RecipeID int    `gorm: "primary_key"`
	Title    string `json:"title"`
	Steps    []Step `json:"steps"` //one-to-many relationship
}

// Quantities know which ingredient they are a part of but ingredients dont need to know the reverse
type Ingredient struct {
	gorm.Model
	Name string `json:"name" gorm:"primary_key"`
}

// Each quantity has one ingredient and one unit of measurement
type Quantity struct {
	gorm.Model

	Ingredient      *Ingredient      //one-to-one relationship
	UnitMeasurement *UnitMeasurement //one-to-one relationship
	Amount          int
	Step            *Step // one-to-one relationship
}

// Each step has many quantities and belongs to only one recipe
type Step struct {
	gorm.Model

	Details    string     `json:"details"`
	Quantities []Quantity `json:"quantities"` // one-to-many relationship
	Recipe     *Recipe    // one-to-one relationship
}

// Measurements are reused by many quantities
type UnitMeasurement struct {
	gorm.Model

	Unit string `gorm:"primary_key`
}
