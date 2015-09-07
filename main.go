package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {

	path := os.Args[1]

	jsonString := os.Args[2]

	var tree map[string]interface{}
	_ := json.Unmarshal([]byte(jsonString), &tree)

	values := make(map[string]interface{})

	for key, value := range tree {
		values[key] = deepGet(value, path)
	}

	outputJSON, _ := json.Marshal(values)

	fmt.Println(string(outputJSON))
}

// function taken from @jwilder/docker-gen
func deepGet(item interface{}, path string) interface{} {
	if path == "" {
		return item
	}

	path = strings.TrimPrefix(path, ".")
	parts := strings.Split(path, ".")
	itemValue := reflect.ValueOf(item)

	if len(parts) > 0 {
		switch itemValue.Kind() {
		case reflect.Struct:
			fieldValue := itemValue.FieldByName(parts[0])
			if fieldValue.IsValid() {
				return deepGet(fieldValue.Interface(), strings.Join(parts[1:], "."))
			}
		case reflect.Map:
			mapValue := itemValue.MapIndex(reflect.ValueOf(parts[0]))
			if mapValue.IsValid() {
				return deepGet(mapValue.Interface(), strings.Join(parts[1:], "."))
			}
		default:
			return nil
		}

	}

	return itemValue.Interface()
}
