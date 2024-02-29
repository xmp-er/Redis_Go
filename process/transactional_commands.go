package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xmp-er/Redis_Go/validatior"
)

func transactional_cmds(s []string) string {
	//we are taking only MULTI from this

	fmt.Println("OK") //inital ok

	var transactional_store []string //storing everything here then will execute at the end

	for {
		var str string
		scanner := bufio.NewScanner(os.Stdin) //validate the input pending
		scanner.Scan()
		str = scanner.Text()

		//checking if the input is correct
		_, err := validatior.Validate_input(str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		st := strings.Split(str, " ")
		if st[0] == "EXEC" {
			break
		} else if st[0] == "DISCARD" {
			return "OK"
		}

		transactional_store = append(transactional_store, str)
		fmt.Println("QUEUED")
	}

	var res string = ""
	for i := 0; i < len(transactional_store); i++ {
		cmd := transactional_store[i]
		st := strings.Split(cmd, " ")
		switch st[0] {
		case "SET", "GET", "DEL":
			res = crud(st)
		case "INCR", "INCRBY":
			res = incr_cmds(st)
		case "MULTI": //recursive MULTI Function maybe
			res = transactional_cmds(st)
		}
		fmt.Println(i+1, ")", res)
	}
	return ""
}
