package main

import (
	"delivery-app/calculation"
	"delivery-app/input"
	"fmt"
)

func main() {
	var userOptions = []string{"Calculate delivery cost", "Calculate the delivery time", "Add offer code", "Show current offers"}
	intIp := input.IntInput{}
	for {
		fmt.Print("\n\n *************  Welcome to Kikis Delivery app ************* \n\n")
		fmt.Println("Please choose from below options & enter the number")

		for i := 0; i < len(userOptions); i++ {
			fmt.Printf("%d - %s \n", i+1, userOptions[i])
		}

		option := intIp.GetIntInput("")
		switch option {
		case 1:
			calculation.GetDeliveryPrice()
		case 2:
			calculation.GetDeliveryTime()
		case 3:
			calculation.AddNewOfferCode()
		case 4:
			calculation.GetOfferCodes()
		default:
			fmt.Println("Please enter a valid option")
		}
	}
}
