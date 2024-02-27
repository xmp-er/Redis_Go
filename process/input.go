package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Map = make(map[string]string)

func main() {
	for {
		var str string = ""
		scanner := bufio.NewScanner(os.Stdin) //validate the input pending
		scanner.Scan()
		str = scanner.Text()
		res := crud(strings.Split(str, " "))
		fmt.Println(res)
	}
}
