package main

import "fmt"

func main(){
	var n int

	fmt.Scanln(&n)
	for n!=0 {
		printSeries(n)
		fmt.Scanln(&n)
	}	
}

func printSeries(n int){
	fmt.Printf("%d ", n)
	max := n
	total := 1
	for n!=1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
		findMaxAndTotal(n, &max, &total)
	}
	fmt.Printf("%d %d\n", max, total)
}

func findMaxAndTotal(n int, max, total *int){
	if n>*max {
		*max = n
	}
	*total++
}