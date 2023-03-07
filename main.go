package main

import (
	"CalorieCounter/data"
	"CalorieCounter/food"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Please log in:")
	fmt.Println("Enter username:")
	username, _ := readLine()
	fmt.Println("Enter password:")
	password, _ := readLine()
	data.Login(username, password)
	fmt.Println("Welcome To The CLI Calorie Counter")
	fmt.Printf("The date is %s.\n", time.Now())

	fmt.Printf("Welcome %s, let's see how we're doing then!\n ", username)

MAINLOOP:
	for {
		fmt.Println("\nPlease select an option:")
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
			fmt.Println("Remove the item by it's number:")
			food.ViewFoodAndCals()
			result, _ := readLine()
			food.Remove(result)

		case "5":
			data.Login("test", "test")
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

func readLine() (line string, err error) {

	rd := bufio.NewReader(os.Stdin)
	line, err = rd.ReadString('\n')
	line = strings.TrimSpace(line)

	if err != nil {
		return "", fmt.Errorf("failed to read line on ReadLine(), Error:%w", err)
	}

	return line, err
}

func readFloat64() (value float64, err error) {

	rd := bufio.NewReader(os.Stdin)

	result, err := rd.ReadString('\n')
	result = strings.TrimSpace(result)

	if err != nil {
		return 0.00, fmt.Errorf("could not read string: %w", err)
	}

	value, err = strconv.ParseFloat(result, 64)

	if err != nil {
		return 0.00, fmt.Errorf("could not read float amount: %w", err)
	}

	return value, nil
}
