package main

import (
	"fmt"
	"errors"
)

var roman_to_int_map = map[string]int {
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// All known pairs where 1st char subtracts from 2nd char
var known_special_pairs = [6]string{"IV", "IX", "XL", "XC", "CD", "CM"}

func main() {
	// Simple in-order addition 	ex) II = 2
	// Single subtraction			ex) IX = 9
	// More than one subtraction	ex) IXXV = 

	romanToInt("X")
	romanToInt("IV")
	romanToInt("IXIIV")
	romanToInt("XXVII")
}

func known_special_pair(s string) bool {
	//fmt.Printf("Checking KSP %s \n", s)
	for _, p := range known_special_pairs {
		if s == p {
			return true
		}
	}
	return false
}

func romanToInt(s string) int {
	accum := 0
	adding_debug := false
	var err error

	for i := 0; i < len(s); i++ {
		val, ok := roman_to_int_map[string(s[i])]

		// Invalid roman numeral
		if !ok {
			err = errors.New("Failed to convert")
			break
		} else {
			if i+2 > len(s) { 
				// Last char, no need to check for special pairs
				if adding_debug { fmt.Printf("Adding %d \n", val) }
				accum += val
			} else {
				// Multiple chars to check, check for special pair
				if known_special_pair(s[i:i+2]) { 
					val2, ok := roman_to_int_map[string(s[i+1])]

					// If next is valid numeral, update accum and increment i
					if !ok {
						err = errors.New("Failed to convert")
						break
					} else {
						if adding_debug { fmt.Printf("Adding (%d - %d) = %d \n", val2, val, val2-val) }
						accum += (val2 - val)
						i++
					}
				} else {
					if adding_debug { fmt.Printf("Adding %d \n", val) }
					accum += val
				}
			}
		}
	}

	if err != nil {
		fmt.Printf("[%s] - %v \n", s, err)
	} else {
		fmt.Printf("[%s] converts to %d \n", s, accum)
	}
	return accum
}