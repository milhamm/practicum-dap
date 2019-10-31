package main

import "fmt"

func main(){
	var n, nOther, total int
	
	fmt.Scanln(&n)
	res := fact(n)
	fmt.Println(res);
	
	for total < n{
		fmt.Scanln(&nOther)
		total += nOther

		res := fact(nOther)
		fmt.Println(res);
	}
}

func fact(n int) int{
	if n == 1{
		return 1
	}

	return n*fact(n-1)
}
