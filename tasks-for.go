package main

import "fmt"

func For39(a, b int) {
	for i := a; i <= b; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println(i)
		}
	}
}

func For40(a, b int) {
	for i := 1; i <= b-a+1; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println(a + i - 1)
		}
	}
}
