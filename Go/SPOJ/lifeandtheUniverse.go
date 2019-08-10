package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var num = scanner.Text()
		if num == "42" {
			break
		} else {
			fmt.Println(num)
		}
	}
}
