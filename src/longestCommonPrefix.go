package main

import (
	"fmt"
)

func main() {
	array_list1 := [3]string{"flower", "flow", "flight"}
	array_list2 := [3]string{"dog","racecar","car"}
	fmt.Println("Longest common prefix = ", longestCommonPrefix(array_list1[:]))
}

func longestCommonPrefix(strs []string) string {
	debug_prints := false

    // 1st loop to find shortest string in the array, that has to be limiting factor for range
    max_range := 1000
    lcp := ""
    cur_str := ""
    fmt.Printf("[ ")
    for _, s := range strs {
    	fmt.Printf(" %s,", s)
    	if len(s) < max_range {
    		max_range = len(s)
    	}
    }
    fmt.Println(" ]")
    if debug_prints { fmt.Println("Max range = ", max_range) }

    for i := 0; i < max_range; i++ {
    	cur_str = strs[0][0:i+1]
    	if debug_prints { fmt.Println("cur_str = ", cur_str) }
    	for _, s := range strs {
    		if debug_prints {  fmt.Printf("Comparing %s %s \n", cur_str, s[0:i+1]) }
    		if cur_str != s[0:i+1] {
    			if debug_prints { fmt.Printf("BREAK! \n") }
    			return lcp
    		}
    	}
    	lcp = cur_str
    }
    return lcp
}