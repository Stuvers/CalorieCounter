package main

import (
	"CalorieCounter/food"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Welcome To The CLI Calorie Counter")
	fmt.Printf("The date is %s.\n", time.Now())

	fmt.Println("Please enter your name:")
	name, _ := readLine()

	fmt.Printf("Welcome %s, let's see how we're doing then!\n ", name)

MAINLOOP:
	for {
		fmt.Println("Please select an option:")
		fmt.Println("1) Today's Calorie Count")
		fmt.Println("2) Add Food")
		fmt.Println("3) List Food Eaten")
		fmt.Println("4) Remove Calories")
		fmt.Println("5) Display All")
		fmt.Println("Q) Quit Program")

		choice, _ := readLine()

		switch choice {
		case "1":
			fmt.Printf("%10.2f\n", food.ViewCalories())
		case "2":

			err := addItem()

			if err != nil {
				fmt.Println(err)
			}
		case "3":
			food.ViewFoodAndCals()
		case "4":
			// do one more thing
		case "5":
			// Surely it's the last thing
		case "Q":
			fmt.Println("LEAVING. . . ")
			break MAINLOOP
		}
	}
}

func addItem() error {

	var calories float64
	var err error

	// Enter the food item name
	fmt.Println("Please enter an item name:")
	itemName, err := readLine()

	if err != nil {
		return fmt.Errorf("failed to record \"itemName\", error:%w", err)
	}

mainloop:
	for {
		// Enter the calories of the item
		fmt.Println("Please enter the calorie amount:")
		calories, err = readFloat64()

		if err != nil {
			return fmt.Errorf("failed to record calorie amount, error:%w", err)
		} else {
			break mainloop
		}
	}

	food.Add(itemName, calories)

	return nil
}

func readLine() (string, error) {

	rd := bufio.NewReader(os.Stdin)
	text, err := rd.ReadString('\n')
	text = strings.ToUpper(strings.TrimSpace(text))

	if err != nil {
		return "", fmt.Errorf("failed to read line on ReadLine(), Error:%w", err)
	}

	return text, nil
}

func readFloat64() (float64, error) {

	floatVal float64
	rd := bufio.NewReader(os.Stdin)

	result, err := rd.ReadString('\n')
	result = strings.TrimSpace(result)

	if err != nil {
		return 0.00, fmt.Errorf("could not read string: %w", err)
	}

	floatVal, err = strconv.ParseFloat(result, 64)

	if err != nil {
		return 0.00, fmt.Errorf("could not read float amount: %w", err)
	}

	return floatVal, nil
}
