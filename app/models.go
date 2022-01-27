package main

import "gorm.io/gorm"

// Each recipe has many steps
type Recipe struct {
	gorm.Model
	Title string
	Steps []Step
}

// Quantities know which ingredient they are a part of but ingredients dont need to know the reverse
type Ingredient struct {
	gorm.Model
	Name string
}

// Each quantity has one ingredient and one unit of measurement
type Quantity struct {
	gorm.Model

	IngredientID      int
	UnitMeasurementID int
	Amount            int
	StepID            int
}

// Each step has many quantities and belongs to only one recipe
type Step struct {
	gorm.Model

	Details    string
	Quantities []Quantity
	RecipeID   int
	Recipe     Recipe
}

// Measurements are reused by many quantities
type UnitMeasurement struct {
	gorm.Model

	Unit string
}
