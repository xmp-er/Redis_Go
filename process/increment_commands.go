package main

import (
	"fmt"
	"strconv"

	"github.com/xmp-er/Redis_Go/validatior"
)

func incr_cmds(s []string) string {

	//performing operations based on command
	switch s[0] {
	case "INCR":
		k := s[1]
		//checking if the value exists,if does not then we make a key by value 1
		if _, ok := Map[k]; !ok {
			op_set([]string{"SET", k, "1"})
		}
		//checking if the key's value is Integer
		err := validatior.Is_Val_Integer(k)
		if err != nil {
			return err.Error()
		}
		//value is integer so, incrementing it by 1 if map has value as integer
		v, err := strconv.Atoi(Map[k])
		if err != nil {
			return "(error) ERR value is not an integer or out of range"
		}
		Map[k] = fmt.Sprintf("%d", v+1)
		return fmt.Sprintf("(integer) %s", Map[k])
	case "INCRBY":
		k := s[1]
		v := s[2]
		//checking if the second value is integer or not
		err := validatior.Is_Val_Integer(v)
		if err != nil {
			return err.Error()
		}
		//if key dne then set value provided
		if _, ok := Map[k]; !ok {
			op_set([]string{"SET", k, v})
		}
		//if key exists then if value is integer then proceeding else error
		temp, err := strconv.Atoi(Map[k])
		if err != nil {
			return "(error) ERR value is not an integer or out of range"
		}
		increment_value, _ := strconv.Atoi(v)
		Map[k] = fmt.Sprintf("%d", temp+increment_value)
		return fmt.Sprintf("(integer) %s", Map[k])
	default:
		return "(error) ERR syntax error"
	}
}
