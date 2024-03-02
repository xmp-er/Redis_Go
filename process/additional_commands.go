package main

import (
	"fmt"
	"strconv"
)

func additional_commands(st []string, Map map[string]string, Backup_Map map[string]string) string {
	//if our map is empty then the final command resulted in empty value and we can return nil
	if len(Map) == 0 {
		return "(nil)"
	}
	// there are values, taking them as final and returning
	var res string = ""
	index := 0
	for k, v := range Map {
		index++
		_, err := strconv.Atoi(v) // if the key is not a integer then not considered
		if err != nil {
			continue
		}
		if Backup_Map[k] == Map[k] { // since we are only adding the key once in BackupMap, if the key was incremented then BackupMap and Map will have different vals for same key
			continue
		}
		fmt.Println(index)
		if index == len(Map)-1 { //avoiding formatting error
			res += (fmt.Sprintf("%s %s %s", "SET", k, v))
		} else {
			res += (fmt.Sprintf("%s %s %s\n", "SET", k, v))
		}
	}
	return res
}
