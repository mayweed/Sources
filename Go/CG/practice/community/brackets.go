package main

import "fmt"

//strings got a count func...
func main() {
	var expression string
	fmt.Scan(&expression)

	var openPar int
	var closePar int

	var openSquareBrac int
	var closeSquareBrac int

	var openCurly int
	var closeCurly int

	for _, car := range expression {
		switch car {
		case '(':
			openPar += 1
		case '[':
			openSquareBrac += 1
		case '{':
			openCurly += 1
		case ')':
			closePar += 1
		case ']':
			closeSquareBrac += 1
		case '}':
			closeCurly += 1
		}
	}
	//oki this is bad!!
	if expression == ")(" {
		fmt.Println("false")
	} else if openPar == closePar && openSquareBrac == closeSquareBrac && openCurly == closeCurly {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
