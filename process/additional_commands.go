package main

import "fmt"

func additional_commands(st []string) string {
	//if our map is empty then the final command resulted in empty value and we can return nil
	if len(Map) == 0 {
		return "(nil)"
	}
	// there are values, taking them as final and returning
	var res string = ""
	index := 0
	for k, v := range Map {
		index++
		if index == len(Map) { //avoiding formatting error
			res += (fmt.Sprintf("%s %s %s", "SET", k, v))
			continue
		}
		res += (fmt.Sprintf("%s %s %s\n", "SET", k, v))
	}
	return res
}
