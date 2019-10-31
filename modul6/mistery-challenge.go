package main

import "fmt"

func main(){
	var n int

	fmt.Scan(&n)
	printSeries(n)
}

func printSeries(n int){
	var initialN int = n
	for n!=1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	if n == 1 {
		fmt.Printf("%d is stopped", initialN)
	}
}