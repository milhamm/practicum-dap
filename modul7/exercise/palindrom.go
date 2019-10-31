package main

import "fmt"

const DOT string = "."
const NMAX int = 127
type table [NMAX]string

func main(){
	var tab table
	var m int

	fillArray(&tab, &m)
	reverseArray(&tab, m)
	printArray(tab, m)

	fmt.Printf("Name \t: Muhammad Ilham Mubarak\n")
	fmt.Printf("SID \t: 1301194276\n")
	
}

func fillArray(t *table, n *int){
	var char string
	*n = 0
	fmt.Scan(&char)
	for char != DOT && *n<=NMAX {
		t[*n] = char
		fmt.Scan(&char)
		*n++
	}
}

func reverseArray(t *table, n int){
	var s = 0
	var e = n-1
	for s < e {
		var temp = (*t)[s]
		(*t)[s] = (*t)[e]
		(*t)[e] = temp
		s++
		e--
	}
}

func printArray(t table, n int){
	fmt.Printf("Palindrom? %t\n", palindrom(t, n))
}

func palindrom(t table, n int) bool{
	var initialArr = t
	reverseArray(&t, n)
	return initialArr == t
}

