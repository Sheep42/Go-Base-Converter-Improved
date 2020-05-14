package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	askConv := true
	askBase := true
	askNum := true

	baseInt := 0
	numInt := 0

	var err error
	var choice string
	var num string

	fmt.Print("*** Go Base Converter ***")

	for {
		//Get the conversion type
		if(askConv) {
			fmt.Print("\n\n(D)ec to Base")
			fmt.Print("\n(B)ase to Dec")

			fmt.Print("\n-> ")
			
			choice, _ = reader.ReadString('\n')
			choice = strings.Replace(choice, "\n", "", -1)

			switch choice {
				case "D": fallthrough
				case "d":
					fmt.Print("\n\nConverting Decimal to Base")
				case "B": fallthrough
				case "b":
					fmt.Print("\n\nConverting Base to Decimal")

				default:
					fmt.Print("\n\nThat was not a choice, please try again")
					continue
			}
		} else {
			askConv = true
		}

		//Get the base
		if(askBase) {
			fmt.Print("\nPlease enter base: ")

			base, _ := reader.ReadString('\n')
			base = strings.Replace(base, "\n", "", -1)

			baseInt, err = strconv.Atoi(base)

			//Catch errors
			if err != nil {
				fmt.Print("\nBase must be an integer\n\n")
				askConv = false
				continue
			}

			if(baseInt < 1 || baseInt > 16) {
				fmt.Print("\nBase must be between 1 and 16\n\n")

				askConv = false

				continue
			}
		} else {
			askBase = true
		}

		//Get the number
		if(askNum) {
			fmt.Print("\nPlease enter number: ")

			num, _ = reader.ReadString('\n')
			num = strings.Replace(num, "\n", "", -1)

			numInt, err = strconv.Atoi(num)

			//Catch errors
			if err != nil && baseInt <= 10 {
				fmt.Print("\nNumber must be an integer\n\n")

				askConv = false
				askBase = false

				continue
			}
		} else {
			askNum = true
		}

		if choice == "D" || choice == "d" {
			result := convertDecToBase(numInt, baseInt)

			fmt.Printf("%d in base %d is %s", numInt, baseInt, result)
		} else if choice == "B" || choice == "b" {
			result := convertBaseToDec(num, baseInt)

			if result >= 0 {
				fmt.Printf("%s (base %d) in decimal is %d", num, baseInt, result)
			}
		}
	}
}

/**
*	Let dec = the number to convert
*	Let base = base to convert to
*	
*	While dec > 0:
*		Let quot = the quotient of dec / base ignoring remainder
*		Let rem = remainder of dec / base
*		Prepend remainder to resulting number (converting to A-F for base 11-16)
*		Let new dec = quot
**/
func convertDecToBase(dec, base int) string {
	resultSlice := []string{} //Declare empty slice
	numberRep := []string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}

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

/**
*	 Let num = string rep of starting number
*	 Let base = base to convert from
*	 Let numLen = length of num
*	 
*	 While numLen > 0:
*		first = leftmost char in num
*		num = num excluding first
*		check if base > 10 and make any necessary conversions
*		convert first to an int
*		decrement numLen
*		add first * (base^numLen) to result
**/
func convertBaseToDec(num string, base int) int {
	result := 0

	numLen := len(num)

	//Read the num char by char
	for numLen > 0 {
		first := num[0:1]
		num = num[1:numLen]

		if base > 10 {
			switch first {
				case "a": fallthrough
				case "A": first = "10"
				case "b": fallthrough
				case "B": first = "11"
				case "c": fallthrough
				case "C": first = "12"
				case "d": fallthrough
				case "D": first = "13"
				case "e": fallthrough
				case "E": first = "14"
				case "f": fallthrough
				case "F": first = "15"
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
    p := 1

    for b > 0 {
        if b & 1 != 0 {
	        p *= a
        }

        b >>= 1
        a *= a
    }

    return p
}