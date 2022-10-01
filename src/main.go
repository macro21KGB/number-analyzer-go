package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(x, y int) int {

	lcm := 1

	if x > y {
		lcm = x
	} else {
		lcm = y
	}
	for {
		if lcm%x == 0 && lcm%y == 0 {
			return lcm
		}
		lcm++
	}
}

func ConvertToExadecimal(val int) string {
	outputtedString := ""

	for val > 0 {

		currentReminder := val % 16

		if currentReminder <= 10 {
			outputtedString = strconv.Itoa(currentReminder) + outputtedString
		} else if currentReminder == 11 {
			outputtedString = "A" + outputtedString
		} else if currentReminder == 12 {
			outputtedString = "B" + outputtedString
		} else if currentReminder == 13 {
			outputtedString = "C" + outputtedString
		} else if currentReminder == 14 {
			outputtedString = "D" + outputtedString
		} else if currentReminder == 15 {
			outputtedString = "E" + outputtedString
		} else if currentReminder == 16 {
			outputtedString = "F" + outputtedString
		}

		// go next iteration
		val = val / 16

	}
	return outputtedString
}

func ConvertNumberToBase(val int, base int) string {

	outputtedString := ""

	for val > 0 {

		reminderOfDivision := val % base

		outputtedString = strconv.Itoa(reminderOfDivision) + outputtedString

		// go next iteration
		val = val / base

	}
	return outputtedString
}

func Initialize() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func ExecuteBaseConversionForValue(value string) {

	convertedNumber, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(Red + "ERROR: Insert a vaild number" + Reset)
		return
	}

	fmt.Println("---- " + value + " ----")
	fmt.Println("Binary: " + Blue + ConvertNumberToBase(convertedNumber, 2) + Reset)
	fmt.Println("Otto: " + Green + ConvertNumberToBase(convertedNumber, 8) + Reset)
	fmt.Println("Excadecimal: " + Yellow + ConvertToExadecimal(convertedNumber) + Reset)
	fmt.Println("----------")
}

func ExecuteMultipleNumbersOperation(val1, val2 string) {

	convertedNumber1, err := strconv.Atoi(val1)
	if err != nil {
		fmt.Println(Red + "ERROR: Insert a vaild number" + Reset)
		return
	}

	convertedNumber2, err := strconv.Atoi(val2)
	if err != nil {
		fmt.Println(Red + "ERROR: Insert a vaild number" + Reset)
		return
	}

	fmt.Printf("MCD: %s%d%s\n", Green, GCD(convertedNumber1, convertedNumber2), Reset)
	fmt.Printf("mcm: %s%d%s\n", Red, LCM(convertedNumber1, convertedNumber2), Reset)

}

func main() {

	Initialize()

	nameOfTheProgram := os.Args[0]

	if len(os.Args) == 1 {
		fmt.Printf("%sUsage: %s <number> Ex: %s 10 %s\n", Yellow, nameOfTheProgram, nameOfTheProgram, Reset)
		fmt.Printf("%sUsage: %s <number1> <number2> Ex: %s 10 5, it will also calculate GCD and LCM %s\n", Yellow, nameOfTheProgram, nameOfTheProgram, Reset)
		return
	}

	if len(os.Args) == 3 {
		selectedNumber1 := os.Args[1]
		selectedNumber2 := os.Args[2]

		ExecuteBaseConversionForValue(selectedNumber1)
		ExecuteBaseConversionForValue(selectedNumber2)
		ExecuteMultipleNumbersOperation(selectedNumber1, selectedNumber2)

		return
	}

	selectedNumber := os.Args[1]
	ExecuteBaseConversionForValue(selectedNumber)

}