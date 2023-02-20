package food

import (
	"fmt"
	"time"
)

var foodItems = map[string]foodItem{
	"Chocolate": {248, time.Now().String()},
	"Apple":     {48, time.Now().String()},
}

type myInfo struct {
	name     string
	tgWeight float64
	calTg    int
}

type foodItem struct {
	calories  float64
	dateAdded string
}

func add(name string, calories float64) {
	foodItems[name] = foodItem{calories: calories, dateAdded: time.Now().String()}
}

func remove(name string) {
	delete(foodItems, name)
}

func Add(name string, calories float64) {
	// Do something at some point
	add(name, calories)
}

func Remove(name string) {
	remove(name)
}

func ViewCalories() float64 {

	var totCals float64 = 0.00
	for _, fi := range foodItems {
		totCals += fi.calories
	}

	return totCals
}

func ViewFoodAndCals() {
	for i, fItem := range foodItems {
		fmt.Printf("\t%s\t%.2f\t%s\n", i, fItem.calories, fItem.dateAdded)
	}
}
