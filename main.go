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

	var number_to_convert string
	var num_as_slice []int
	var num_as_int int
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
		if "" == number_to_convert {

			fmt.Print("\nPlease enter number to convert: ")

			if number_to_convert, err = getInput(); err != nil {
				fmt.Println(err)
				continue
			}

			if num_as_slice, err = getNumberAsSlice(number_to_convert); err != nil {
				fmt.Println(err)
				number_to_convert = ""
				continue
			}

			if err = validateNumber(num_as_slice, base, choice); err != nil {
				fmt.Println(err)
				number_to_convert = ""
				continue
			}

		}

		if "d" == choice {

			if num_as_int, err = strconv.Atoi(number_to_convert); err != nil {
				fmt.Print(err)
				number_to_convert = ""
				continue
			}

			result := convertDecToBase(num_as_int, base)
			fmt.Printf("%d in base %d is %s", num_as_int, base, result)

		} else if "b" == choice {

			result := convertBaseToDec(num_as_slice, base)
			fmt.Printf("%s (base %d) in decimal is %d", number_to_convert, base, result)

		}

		choice = ""
		base = 0
		number_to_convert = ""

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
		num_length--

	}

	return num_as_slice, nil

}

func validateNumber(num_slice []int, base int, choice string) error {

	upper_limit := base

	if 1 == base {
		upper_limit = 2
	} 

	if "d" == choice {
		upper_limit = 10
	}

	for _, val := range num_slice {

		if val >= upper_limit {
			return errors.ThrowInputError(strconv.Itoa(val), fmt.Sprintf("is greater than or equal to base: %d", upper_limit))
		} 

	}

	return nil

}

func convertDecToBase(decimal, base int) string {

	result_as_slice := []string{}
	number_glyphs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	if 0 == decimal {
		return "0"
	} 

	for decimal > 0 {

		if base > 1 {

			quotient := decimal / base
			remainder := decimal % base

			//Prepend remainder to result
			remainder_as_slice := []string{number_glyphs[remainder]}
			result_as_slice = append(remainder_as_slice, result_as_slice...)

			decimal = quotient

		} else {

			//Handle unary (lol)
			result_as_slice = append(result_as_slice, "1")
			decimal--

		}

	}

	result := strings.Join(result_as_slice, "")

	return result

}

func convertBaseToDec(num_as_slice []int, base int) int {

	result := 0
	num_length := len(num_as_slice)

	if base > 1 {
		
		for _, val := range num_as_slice {

			num_length--
			result += (val * (raise(base, num_length)))

		}

	} else {

		for i := 0; i < num_length; i++ {
			result++;
		}

	}

	return result

}

//Integer exponentation
func raise(base, power int) int {

	result := 1

	for power > 0 {

		if power&1 != 0 {
			result *= base
		}

		power >>= 1
		base *= base

	}

	return result

}
