package main

import (
	"fmt"
)

func countChars(str string) map[string]int {
	hash := make(map[string]int)
	for _, s := range str {
		j := hash[string(s)]

		if j == 0 {
			hash[string(s)] = 1
		} else {
			hash[string(s)] = j + 1
		}
	}
	return hash
}

func areAnagrams(s1 string, s2 string) bool {
	if len(s1) == len(s2) {
		if s1 == s2 {
			fmt.Println("They are the same!")
			return false
		}
		// return reflect.DeepEqual(countChars(s1), countChars(s2))
		var res bool = true
		s1_ana := countChars(s1)
		s2_ana := countChars(s2)
		for s1k, s1v := range s1_ana {
			res = (s2_ana[s1k] == s1v) && res
		}
		return res
	}
	return false
}

func main() {
	str1 := ""
	str2 := ""
	fmt.Printf("Enter string 1: ")
	fmt.Scanf("%s", &str1)
	fmt.Printf("Enter string 2: ")
	fmt.Scanf("%s", &str2)
	fmt.Println("Is anagram: ", areAnagrams(str1, str2))
}
