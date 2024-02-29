package main

import (
	"fmt"
	"strconv"

	"github.com/xmp-er/Redis_Go/validatior"
)

func incr_cmds(s []string) string {
	k := s[1]
	//checking if the value exists,if does not then we make a key by value 1
	if _, ok := Map[k]; !ok {
		op_set([]string{"SET", "k", "1"})
	}
	//checking if the key's value is Integer
	err := validatior.Is_Val_Integer(k)
	if err != nil {
		return err.Error()
	}
	//value is integer so, incrementing it by 1
	v, _ := strconv.Atoi(k)
	Map[k] = fmt.Sprintf("%d", v+1)
	return fmt.Sprintf("(integer) %s", Map[k])
}
