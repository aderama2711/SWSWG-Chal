package main

import (
	"fmt"
)

func main() {
	str := "selamat malam"

	split, count := calcword(str)

	for _, v := range split {
		fmt.Println(v)
	}

	fmt.Println(count)
}

func calcword(str string) ([]string, map[string]int) {

	split := []string{}
	count := map[string]int{}

	for _, v := range str {
		split = append(split, string(v))
		_, exist := count[string(v)]
		if exist {
			count[string(v)]++
		} else {
			count[string(v)] = 1
		}
	}

	return split, count
}
