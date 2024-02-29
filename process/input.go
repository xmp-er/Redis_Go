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
			continue
		}

		//splitting the string into array
		st := strings.Split(str, " ")

		//string holding our final result
		var res string = ""

		//if the command is GET,SET or DEL, we handle it via the crud_commands handler
		switch st[0] {
		case "SET", "GET", "DEL":
			res = crud(st)
		case "INCR", "INCRBY":
			res = incr_cmds(st)
		}

		fmt.Println(res)
	}
}
