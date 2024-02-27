package process

import (
	"fmt"
	"strconv"
	"strings"
)

//1 stands for successful op, 0/nil for failed op

func crud(str []string) int { //expectation is that the incoming command must be SET,GET, OR DEL
	if str[0] == "SET" {
		op_set(str)
	} else if str[0] == "GET" {
		op_get(str)
	} else if str[0] == "DEL" {
		op_del(str)
	}
}

func op_set(str []string) int { //SET operation implementation, assumption all input correct
	k := str[1]
	v := str[2]
	temp_v, err := strconv.ParseFloat(v, 64)
	if err != nil {
		//it means we have a float64 and we can assign it to Map_int
		Map_int[k] = temp_v
		return 1
	}
	//now the remaining can only be the string value so adding it in Map_string
	Map_string[k] = strings.Trim(v, "\"")
	return 1
}

func op_get(str []string) string { //GET implementation, assuming the input to be correct
	k := str[1]
	if _, ok := Map_string[k]; !ok {
		return fmt.Sprintf("%f", Map_int[k])
	}
	return Map_string[k]
}
