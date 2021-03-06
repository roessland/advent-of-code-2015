package main

import "fmt"
import "strconv"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

func main() {
	polymer := []rune{}

	buf, _ := ioutil.ReadFile("input.txt")
	for _, char := range string(buf) {
		if char == '\n' {
			continue
		}
		polymer = append(polymer, char)
	}
	poly := strings.Trim(string(buf), " \n\r")
	_ = poly

	for {
		for c := 'a'; c <= 'z'; c++ {
			C := c + 'A' - 'a'
			poly = strings.Replace(poly, fmt.Sprintf("%c%c", c, C), "", -1)
			poly = strings.Replace(poly, fmt.Sprintf("%c%c", C, c), "", -1)
		}
		fmt.Println(len(poly))
	}
}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
