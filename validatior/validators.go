package validatior

import (
	"errors"
	"strconv"
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
	case "GET", "DEL", "INCR":
		if !Is_Two_Args(s) {
			return false, errors.New("(error) ERR syntax error")
		}
	case "SET", "INCRBY":
		if !Is_set_valid(s) {
			return false, errors.New("(error) ERR syntax error")
		}
	case "MULTI", "EXEC", "DISCARD", "COMPACT", "DISCONNECT":
		if !Is_One_Args(s) {
			return false, errors.New("(error) ERR syntax error")
		}
	case "SELECT":
		if !Is_Two_Args(s) {
			return false, errors.New("(error) ERR syntax error")
		}
		if !is_DB_Range_Integer(s) {
			return false, errors.New("(error) ERR value is not an integer or out of range")
		}
		if !is_DB_Range_Valid(s) {
			return false, errors.New("(error) ERR DB index is out of range")
		}
	}

	return true, nil
}

func Is_Valid_Command(str string) bool { //checks if the command part is valid or not
	switch str {
	case "GET", "SET", "DEL", "INCR", "INCRBY", "MULTI", "EXEC", "DISCARD", "COMPACT", "DISCONNECT", "SELECT": //if command is of type GET,SET or DEl,INCR,INCRBY,MULTI,EXEC,DISCARD and others valid else not
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
	if len(str) > 3 {
		if str[2][0] != '"' || str[len(str)-1][len(l)-1] != '"' {
			return false
		}
	}
	return true
}

func Is_Val_Integer(v string) error {
	_, err := strconv.Atoi(v)
	if err != nil {
		return errors.New("(error) ERR value is not an integer or out of range")
	}
	return nil
}

func Is_One_Args(str []string) bool {
	return len(str) == 1
}

func Is_Valid_Port(s string) bool {
	if s[0] != ':' || len(s) != 5 {
		return false
	}
	_, err := strconv.Atoi(s[1:])
	return err == nil
}

func is_DB_Range_Integer(s []string) bool {
	_, err := strconv.Atoi(s[1])
	if err != nil {
		return false
	}
	return true
}

func is_DB_Range_Valid(s []string) bool {
	v, _ := strconv.Atoi(s[1])

	if v < 0 || v > 15 {
		return false
	}
	return true
}
