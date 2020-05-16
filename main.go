package main

import (
	"Go-Base-Converter-Improved/errors"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const MIN_BASE = 1
	const MAX_BASE = 16
	valid_opts := []string{"d", "b", "q"}

	var num string
	var num_as_slice []int
	var base int
	var choice string
	var err error

	options := map[string]string{
		"d": "Converting decimal to base",
		"b": "Converting base to decimal",
	}

	fmt.Print("*** Go Base Converter ***")

	for {

		// Get the option
		if "" == choice {

			fmt.Print("\n\n(D)ecimal to Base")
			fmt.Print("\n(B)ase to Decimal")
			fmt.Print("\n(Q)uit")
			fmt.Print("\n-> ")

			choice, err = getStringInput(valid_opts)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if "q" == choice {
				fmt.Print("\nBye")
				break
			}

			fmt.Print(options[choice])

		}

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
		if "" == num {

			fmt.Print("\nPlease enter number to convert: ")

			if num, err = getInput(); err != nil {
				fmt.Println(err)
				continue
			}

			if num_as_slice, err = getNumberAsSlice(num); err != nil {
				fmt.Println(err)
				continue
			}

			if err = validateNumber(num_as_slice, base); err != nil {
				fmt.Println(err)
				continue
			}

		}

	}
}

func getStringInput(valid_opts []string) (string, error) {

	choice, err := getInput()

	_, found := inSlice(choice, valid_opts)

	if !found {
		return "", errors.ThrowInputError(choice, "is not a choice")
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

func processOptions(choice string) {

	// if "d" == choice {

	// 	result := convertDecToBase(numInt, base)

	// 	fmt.Printf("%d in base %d is %s", numInt, base, result)

	// } else if "b" == choice {

	// 	result := convertBaseToDec(num, base)

	// 	if result >= 0 {
	// 		fmt.Printf("%s (base %d) in decimal is %d", num, base, result)
	// 	}

	// }

}

func getNumberInput(min int, max int) (int, error) {

	choice, err := getInput()

	if err != nil {
		return 0, err
	}

	choice_as_num, err := strconv.Atoi(choice)

	if err != nil {
		return 0, errors.ThrowInputError(choice, "is not a number")
	}

	if choice_as_num < min || choice_as_num > max {
		return 0, errors.ThrowInputError(choice, fmt.Sprintf(" is not between %d and %d", min, max))
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

func validateNumber(num_slice []int, base int) error {

	for _, val := range num_slice {

		if val >= base {
			return errors.ThrowInputError(strconv.Itoa(val), fmt.Sprintf("is greater than or equal to base %d", base))
		}

	}

	return nil

}

func getNumberAsSlice(num string) ([]int, error) {

	var err error
	var ok bool

	first_as_int := 0

	char_mappings := map[string]int{
		"a": 10,
		"b": 11,
		"c": 12,
		"d": 13,
		"e": 14,
		"f": 15,
	}

	num_as_slice := []int{}
	num_length := len(num)

	for num_length > 0 {

		first := num[0:1]
		num = num[1:num_length]

		first_as_int, err = strconv.Atoi(first)

		if err != nil {

			first_as_int, ok = char_mappings[first]

			if !ok {
				return []int{}, errors.ThrowInputError(first, "is not a valid value")
			}

		}

		num_as_slice = append(num_as_slice, first_as_int)

	}

	return num_as_slice, nil

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
