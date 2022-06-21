package main

import "fmt"

func main() {
	input := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}
	city := make(map[string]int)

	for _, innerObject := range input {

		for i, k := range innerObject {
			if i == 0 {
				city[k] = 0
			} else {

				if _, ok := city[k]; ok {
					delete(city, k)
				} else {
					city[k] = 1
				}
			}
		}

	}

	result := ""

	for key, value := range city {
		if value == 1 {
			result = key
			break
		}
	}

	fmt.Println(result)
}
