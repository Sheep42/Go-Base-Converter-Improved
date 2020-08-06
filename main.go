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

	const MinBase = 1
	const MaxBase = 16
	validOpts := []string{"d", "b", "q"}

	var numberToConvert string
	var numAsSlice []int
	var numAsInt int
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

			choice, err = getStringInput(validOpts)

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

			base, err = getNumberInput(MinBase, MaxBase)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		//Get the number
		if "" == numberToConvert {

			fmt.Print("\nPlease enter number to convert: ")

			if numberToConvert, err = getInput(); err != nil {
				fmt.Println(err)
				continue
			}

			if numAsSlice, err = getNumberAsSlice(numberToConvert); err != nil {
				fmt.Println(err)
				numberToConvert = ""
				continue
			}

			if err = validateNumber(numAsSlice, base, choice); err != nil {
				fmt.Println(err)
				numberToConvert = ""
				continue
			}

		}

		if "d" == choice {

			if numAsInt, err = strconv.Atoi(numberToConvert); err != nil {
				fmt.Print(err)
				numberToConvert = ""
				continue
			}

			result := convertDecToBase(numAsInt, base)
			fmt.Printf("%d in base %d is %s", numAsInt, base, result)

		} else if "b" == choice {

			result := convertBaseToDec(numAsSlice, base)
			fmt.Printf("%s (base %d) in decimal is %d", numberToConvert, base, result)

		}

		choice = ""
		base = 0
		numberToConvert = ""

	}

}

func getStringInput(validOpts []string) (string, error) {

	choice, err := getInput()

	_, found := inSlice(choice, validOpts)

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

	choiceAsNum, err := strconv.Atoi(choice)

	if err != nil {
		return 0, errors.ThrowInputError(choice, "is not a number")
	}

	if choiceAsNum < min || choiceAsNum > max {
		return 0, errors.ThrowInputError(choice, fmt.Sprintf(" is not between %d and %d", min, max))
	}

	return choiceAsNum, err

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

	firstAsInt := 0

	charMappings := map[string]int{
		"a": 10,
		"b": 11,
		"c": 12,
		"d": 13,
		"e": 14,
		"f": 15,
	}

	numAsSlice := []int{}
	numLength := len(num)

	for numLength > 0 {

		first := num[0:1]
		num = num[1:numLength]

		firstAsInt, err = strconv.Atoi(first)

		if err != nil {

			firstAsInt, ok = charMappings[first]

			if !ok {
				return []int{}, errors.ThrowInputError(first, "is not a valid value")
			}

		}

		numAsSlice = append(numAsSlice, firstAsInt)
		numLength--

	}

	return numAsSlice, nil

}

func validateNumber(numSlice []int, base int, choice string) error {

	upperLimit := base

	if 1 == base {
		upperLimit = 2
	}

	if "d" == choice {
		upperLimit = 10
	}

	for _, val := range numSlice {

		if val >= upperLimit {
			return errors.ThrowInputError(strconv.Itoa(val), fmt.Sprintf("is greater than or equal to base: %d", upperLimit))
		}

	}

	return nil

}

func convertDecToBase(decimal, base int) string {

	resultAsSlice := []string{}
	numberGlyphs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	if 0 == decimal {
		return "0"
	}

	for decimal > 0 {

		if base > 1 {

			quotient := decimal / base
			remainder := decimal % base

			//Prepend remainder to result
			remainderAsSlice := []string{numberGlyphs[remainder]}
			resultAsSlice = append(remainderAsSlice, resultAsSlice...)

			decimal = quotient

		} else {

			//Handle unary (lol)
			resultAsSlice = append(resultAsSlice, "1")
			decimal--

		}

	}

	result := strings.Join(resultAsSlice, "")

	return result

}

func convertBaseToDec(numAsSlice []int, base int) int {

	result := 0
	numLength := len(numAsSlice)

	if base > 1 {

		for _, val := range numAsSlice {

			numLength--
			result += (val * (raise(base, numLength)))

		}

	} else {

		for i := 0; i < numLength; i++ {
			result++
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
