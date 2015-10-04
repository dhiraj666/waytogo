//season

package main

import (
	"fmt"
)

func main() {

	var monthNumer int = 4

	switch a := season(monthNumer); {
	case "autumn":
		fmt.Println("season is autumn")
	case "spring":
		fmt.Println("season is spring")
	case "winter":
		fmt.Println("season is winter")
	case "summer":
		fmt.Println("season is summer")

	default:
		fmt.Println("No idea which season is ")
	}

}

func season(monthNumber int) string {
	if monthNumber <= 4 {
		return "winter"
	} else if monthNumber > 4 && monthNumber <= 8 {
		return "summer"
	} else {
		return "spring"
	}

}
