package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	s1 := "cat"
	s2 := "tac"

	if isAnagram(s1, s2) {
		fmt.Println("Is an anagram", s1, s2)
	} else {
		fmt.Println("Is not an Anagram", s1, s2)
	}

}

// brute force method
// uses 2 loops, not nested so O(n) still, specifically O(s+t) since we have to loop through both
// does use the Go standard library to check if maps are equal
func isAnagram(s string, t string) bool {
	//return immediately if the strings aren't the same length
	if len(s) != len(t) {
		return false
	}

	//puts all of the runes of both slices into a map and then compares both maps
	sRune := []rune(s)
	tRune := []rune(t)
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for i := 0; i < len(sRune); i++ {
		sMap[sRune[i]] += 1
		tMap[tRune[i]] += 1
	}

	return reflect.DeepEqual(sMap, tMap)
}

//another way to solve the problem through sorting
//this is slower but more memory efficient
func isAnagram2(s string, t string) bool {
	sRune := []rune(s)
	tRune := []rune(t)
	sort.Slice(sRune, func(i, j int) bool {
		return sRune[i] < sRune[j]
	})
	sort.Slice(tRune, func(i, j int) bool {
		return tRune[i] < tRune[j]
	})

	return reflect.DeepEqual(sRune, tRune)
}
