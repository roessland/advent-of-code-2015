package main

import "fmt"
import "io/ioutil"
import "strings"

func CountVowels(s string) int {
	numvowels := 0
	for _, vowel := range []string{"a", "e", "i", "o", "u"} {
		numvowels += strings.Count(s, vowel)
	}
	return numvowels
}

func TwoInARow(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func NoBadStrings(s string) bool {
	for _, bad := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Count(s, bad) > 0 {
			return false
		}
	}
	return true
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\n")
	nice := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if CountVowels(line) >= 3 && TwoInARow(line) && NoBadStrings(line) {
			nice++
		}
	}
	fmt.Printf("nice: %d\n", nice)
}
