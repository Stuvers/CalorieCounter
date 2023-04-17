package food

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserFood(user string) ([]foodItem, error) {

	foodItems := []foodItem{}

	file, err := os.Open("data/users.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	allowRead := false

	for scanner.Scan() {
		if scanner.Text() == "UserFood" {
			allowRead = true
		} else if allowRead {
			userFdDets := strings.Split(scanner.Text(), ",")

			if userFdDets[0] == user {
				cals, _ := strconv.ParseFloat(userFdDets[2], 64)
				foodItems = append(foodItems, foodItem{userFdDets[1], cals, userFdDets[3]})
			}
		}
	}

	return foodItems, err
}
