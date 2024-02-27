package process

import (
	"fmt"
	"strconv"
	"strings"
)

//1 stands for successful op, 0/nil for failed op

func crud(str []string) string { //expectation is that the incoming command must be SET,GET, OR DEL
	var res string = ""
	if str[0] == "SET" {
		res = op_set(str)
	} else if str[0] == "GET" {
		res = op_get(str)
	} else if str[0] == "DEL" {
		res = op_del(str)
	}
	return res
}

func op_set(str []string) string { //SET operation implementation, assumption all input correct
	k := str[1]
	v := str[2]
	temp_v, err := strconv.ParseFloat(v, 64)
	if err != nil {
		//it means we have a float64 and we can assign it to Map_int
		Map_int[k] = temp_v
		return "1"
	}
	//now the remaining can only be the string value so adding it in Map_string
	Map_string[k] = strings.Trim(v, "\"")
	return "1"
}

func op_get(str []string) string { //GET implementation, assuming the input to be correct
	k := str[1]
	if _, ok := Map_string[k]; !ok {
		return fmt.Sprintf("%f", Map_int[k])
	}
	if _, ok := Map_string[k]; !ok {
		return fmt.Sprintf("%f", Map_int[k])
	}
	return "0"
}

func op_del(str []string) string { //DEL implementation, assuming input to be correct
	k := str[1]
	if _, ok := Map_string[k]; ok {
		delete(Map_string, k)
		return "1"
	}
	if _, ok := Map_int[k]; ok {
		delete(Map_int, k)
		return "1"
	}
	return "0"
}
