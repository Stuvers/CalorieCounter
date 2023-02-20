package food

import (
	"time"
)

type myInfo struct {
	name     string
	tgWeight float64
	calTg    int
}

type foodItem struct {
	name      string
	calories  float64
	dateAdded string
}

func add(fi foodItem) {
	foodItems = append(foodItems, fi)
}

func Add(name string, calories float64) {
	// Do something at some point
	add(foodItem{name: name, calories: calories, dateAdded: time.DateTime})
}

func ViewCalories() float64 {

	var totCals float64 = 0.00
	for _, fi := range foodItems {
		totCals += fi.calories
	}

	return totCals
}

func ViewFoodAndCals() string {
	for i, y, := range foodItems {
		i.foodItem.
	}
}
