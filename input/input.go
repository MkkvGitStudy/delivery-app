package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FloatInputGetter interface {
	GetFloatInput(string) float64
}

type FloatInput struct{}

type IntInputGetter interface {
	GetIntInput(string) int
}

type IntInput struct{}

func (FloatInput) GetFloatInput(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid float number.")
			continue
		}
		return value
	}
}

func (IntInput) GetIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		return value
	}
}

func GetStringInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Alphanumeric without any space
func GetCouponCodeInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if regex.MatchString(input) {
			return input
		}
		fmt.Println("Invalid input. Please enter a valid alphanumeric string without spaces.")
	}
}

func GetPositiveFloatValue(prompt string, zeroInclusive bool, getter FloatInputGetter) float64 {
	value := getter.GetFloatInput(prompt)
	if zeroInclusive {
		for value < 0 {
			fmt.Println("Please enter a positive value")
			value = getter.GetFloatInput(prompt)
		}
	} else {
		for value <= 0 {
			fmt.Println("Please enter a positive non zero value")
			value = getter.GetFloatInput(prompt)
		}
	}
	return value
}

func GetPositiveIntValue(prompt string, getter IntInputGetter) int {
	value := getter.GetIntInput(prompt)
	for value <= 0 {
		fmt.Println("Please enter a positive value")
		value = getter.GetIntInput(prompt)
	}
	return value
}
