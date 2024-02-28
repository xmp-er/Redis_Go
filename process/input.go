package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xmp-er/Redis_Go/validatior"
)

var Map = make(map[string]string)

func main() {
	for {
		var str string = ""
		scanner := bufio.NewScanner(os.Stdin) //validate the input pending
		scanner.Scan()
		str = scanner.Text()

		//checking if the input is correct
		_, err := validatior.Validate_input(str)
		if err != nil {
			fmt.Println(err)
		}

		res := crud(strings.Split(str, " "))
		fmt.Println(res)
	}
}
