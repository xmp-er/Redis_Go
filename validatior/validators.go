package validatior

import (
	"errors"
	"strings"
)

func Validate_input(str string) (bool, error) {
	s := strings.Split(str, " ")
	if !Is_Valid_Command(s[0]) { //check if command is valid
		return false, errors.New("(error) ERR unknown command")
	}
	//if SET command and spaces then all the following must be enclosed in Quotes if spaces
	//if GET or DEL command, then there must be only two elements in splitted string
	switch s[0] {
	case "GET", "DEL":
		if !Is_Two_Args(s) {
			return false, errors.New("(error) ERR syntax error")
		}
	case "SET":
		if !Is_set_valid(s) {
			return false, errors.New("(error) ERR syntax error")
		}
	}

	return true, nil
}

func Is_Valid_Command(str string) bool { //checks if the command part is valid or not
	switch str {
	case "GET", "SET", "DEL": //if command is of type GET,SET or DEl, valid else not
		return true
	default:
		return false
	}
}

func Is_Two_Args(str []string) bool {
	return len(str) == 2
}

func Is_set_valid(str []string) bool {
	l := str[len(str)-1]
	if len(str) > 2 {
		if str[3][0] != '"' && str[len(str)-1][len(l)-1] != '"' {
			return false
		}
	}
	return true
}
