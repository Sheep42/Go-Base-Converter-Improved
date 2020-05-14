package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputErr struct {
	input string
	msg   string
}

func (e *inputErr) Error() string {
	return fmt.Sprintf("%s %s", e.input, e.msg)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	askNum := true

	numInt := 0

	const MIN_BASE = 1
	const MAX_BASE = 16
	opts := []string{"d", "b", "q"}

	var num string
	var base int
	var choice string
	var err error

	fmt.Print("*** Go Base Converter ***")

	for {

		// Get the option
		if "" == choice {

			fmt.Print("\n\n(D)ecimal to Base")
			fmt.Print("\n(B)ase to Decimal")

			fmt.Print("\n-> ")

			choice, err = getStringInput(opts)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		// fmt.Print("\n\nConverting Decimal to Base")
		// fmt.Print("\n\nConverting Base to Decimal")

		//Get the base
		if 0 == base {

			fmt.Print("\nPlease enter base between 1 and 16: ")

			base, err = getNumberInput(MIN_BASE, MAX_BASE)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		//Get the number
		if askNum {
			fmt.Print("\nPlease enter number: ")

			num, _ = reader.ReadString('\n')
			num = strings.Replace(num, "\n", "", -1)

			numInt, err = strconv.Atoi(num)

			//Catch errors
			if err != nil && base <= 10 {
				fmt.Print("\nNumber must be an integer\n\n")

				continue
			}
		} else {
			askNum = true
		}

		if choice == "D" || choice == "d" {
			result := convertDecToBase(numInt, base)

			fmt.Printf("%d in base %d is %s", numInt, base, result)
		} else if choice == "B" || choice == "b" {
			result := convertBaseToDec(num, base)

			if result >= 0 {
				fmt.Printf("%s (base %d) in decimal is %d", num, base, result)
			}
		}
	}
}

func getStringInput(valid_opts []string) (string, error) {

	choice, err := getInput()

	_, found := inSlice(choice, valid_opts)

	if !found {
		return "", &inputErr{choice, "is not a choice"}
	}

	return choice, err

}

func inSlice(needle string, haystack []string) (int, bool) {

	for key, item := range haystack {

		if item == needle {
			return key, true
		}

	}

	return -1, false
}

func getNumberInput(min int, max int) (int, error) {

	choice, err := getInput()

	if err != nil {
		return 0, err
	}

	choice_as_num, err := strconv.Atoi(choice)

	if err != nil {
		return 0, &inputErr{choice, "is not a number"}
	}

	if choice_as_num < min || choice_as_num > max {
		return 0, &inputErr{choice, fmt.Sprintf(" is not between %d and %d", min, max)}
	}

	return choice_as_num, err

}

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	choice, err := reader.ReadString('\n')
	choice = strings.Replace(choice, "\n", "", -1)

	choice = strings.ToLower(choice)

	return choice, err
}

func convertDecToBase(dec, base int) string {
	resultSlice := []string{} //Declare empty slice
	numberRep := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	for dec > 0 {
		if base > 1 {
			quot := dec / base
			rem := dec % base

			//Prepend rem to result
			remSlice := []string{numberRep[rem]}
			resultSlice = append(remSlice, resultSlice...)

			dec = quot
		} else {
			//Handle unary (lol)
			resultSlice = append(resultSlice, "1")
			dec--
		}
	}

	result := strings.Join(resultSlice, "")

	return result
}

func convertBaseToDec(num string, base int) int {
	result := 0

	numLen := len(num)

	//Read the num char by char
	for numLen > 0 {
		first := num[0:1]
		num = num[1:numLen]

		if base > 10 {
			switch first {
			case "a":
				fallthrough
			case "A":
				first = "10"
			case "b":
				fallthrough
			case "B":
				first = "11"
			case "c":
				fallthrough
			case "C":
				first = "12"
			case "d":
				fallthrough
			case "D":
				first = "13"
			case "e":
				fallthrough
			case "E":
				first = "14"
			case "f":
				fallthrough
			case "F":
				first = "15"
			default:
			}
		}

		firstInt, err := strconv.Atoi(first)

		//Check errors
		if err != nil {
			fmt.Print("\n\nCould not finish conversion: Number contains an invalid value!\n\n")

			return -1
		}

		if firstInt >= base {
			fmt.Print("\n\nCould not finish conversion: Cannot provide numeric value greater than base!\n\n")

			return -1
		}

		numLen--
		result += (firstInt * (pow(base, numLen)))
	}

	return result
}

//Integer exponentation
func pow(a, b int) int {
	result := 1

	for b > 0 {
		if b&1 != 0 {
			result *= a
		}

		b >>= 1
		a *= a
	}

	return result
}
