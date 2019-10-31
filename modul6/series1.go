package main

import "fmt"

func main(){
	var n int

	fmt.Scan(&n)
	fmt.Printf("%d ", n)
	printSeries(n)
}

func printSeries(n int){
	for n!=1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
		fmt.Printf("%d ", n)	
	}
}