package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("\n=================================")
		fmt.Println("    CLI MULTI-UNIT CONVERTER     ")
		fmt.Println("=================================")
		fmt.Println("1. Temperature Converter")
		fmt.Println("2. Distance Converter")
		fmt.Println("3. Currency Converter")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option (1-4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("❌ Invalid input. Please enter a number between 1 and 4.")
			continue
		}
		switch choice{
			case 1 :
				convertTemperature(reader) 
			case 2 :
				convertDistance(reader)
			case 3 :
				convertCurrency(reader)
			case 4 : 
				fmt.Println("Goodbye! Thanks for using the converter.")
				return
		}
	}
}
func convertTemperature(reader *bufio.Reader){
	fmt.Println("\n--- Temperature Converter ---")
	fmt.Println("A. Celsius to Fahrenheit")
	fmt.Println("B. Fahrenheit to Celsius")
	fmt.Print("Choose conversion (A/B): ")
	option , _ := reader.ReadString('\n')
	fmt.Println(option)
	option = strings.ToUpper(strings.TrimSpace(option))
	fmt.Println(option)
	if option != "A" && option != "B" {
		fmt.Println("❌ Invalid choice.")
		return
	}
	fmt.Print("Enter value to convert: ")
	valstr, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(valstr), 64)
	if err != nil {
		fmt.Println("❌ Invalid numeric value.")
		return
	}
	if option == "A" {
		res := (val * 9 / 5) + 32
		fmt.Printf("🎯 %.2f°C = %.2f°F\n", val, res)
	} else {
		res := (val - 32) * 5 / 9
		fmt.Printf("🎯 %.2f°F = %.2f°C\n", val, res)
	}

}
func convertDistance(reader *bufio.Reader) {
	fmt.Println("\n--- Distance Converter ---")
	fmt.Println("A. Meters to Kilometers")
	fmt.Println("B. Kilometers to Miles")
	fmt.Print("Choose conversion (A/B): ")

	option, _ := reader.ReadString('\n')
	option = strings.ToUpper(strings.TrimSpace(option))

	fmt.Print("Enter value to convert: ")
	valStr, _ := reader.ReadString('\n')
	valStr = strings.TrimSpace(valStr)

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Println("❌ Invalid numeric value.")
		return
	}

	switch option {
	case "A":
		fmt.Printf("🎯 %.2f meters = %.4f km\n", val, val/1000)
	case "B":
		fmt.Printf("🎯 %.2f km = %.4f miles\n", val, val*0.621371)
	default:
		fmt.Println("❌ Invalid choice.")
	}
}

// --- 3. CURRENCY CONVERSION ---
func convertCurrency(reader *bufio.Reader) {
	const usdToInr = 83.50

	fmt.Println("\n--- Currency Converter ---")
	fmt.Println("A. USD to INR")
	fmt.Println("B. INR to USD")
	fmt.Print("Choose conversion (A/B): ")

	option, _ := reader.ReadString('\n')
	option = strings.ToUpper(strings.TrimSpace(option))

	fmt.Print("Enter amount: ")
	valStr, _ := reader.ReadString('\n')
	valStr = strings.TrimSpace(valStr)

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Println("❌ Invalid numeric value.")
		return
	}

	switch option {
	case "A":
		fmt.Printf("🎯 $%.2f USD = ₹%.2f INR\n", val, val*usdToInr)
	case "B":
		fmt.Printf("🎯 ₹%.2f INR = $%.2f USD\n", val, val/usdToInr)
	default:
		fmt.Println("❌ Invalid choice.")
	}
}