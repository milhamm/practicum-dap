package main

import "fmt"

func main(){
	var n, nOther, total, productN int

	productN = 1
	
	fmt.Scanln(&n)
	
	for total < n{
		fmt.Scanln(&nOther)
		
		total += nOther
		
		if total > n {
			nOther = (total - nOther) - 1
		}
		
		productN *= fact(nOther)	
	}

	res := permutation(fact(n),productN)
	fmt.Print(res);
}

func fact(n int) int{
	if n == 1{
		return 1
	}

	return n*fact(n-1)
}

func permutation(n, productN int)int{
	return n / productN
}