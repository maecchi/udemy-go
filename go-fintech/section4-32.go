package main

import (
	"fmt"
)

func question1() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for i, v := range l {
		if i == 0 {
			min = v
			continue
		}
		if min > v {
			min = v
		}
	}
	fmt.Println(min)
}

func question2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	total := 0
	for _, price := range m {
		total += price
	}

	fmt.Printf("%d円\n", total)
}

func main() {
	question1()
	question2()

}
