package main

import "fmt"

const N = 7

type RecType struct{
	f1 string
	f2 int
	f3 float64
}

type ArrType [N]RecType

func main(){
	var isinya ArrType = ArrType{
		RecType{"Ancok",2,3.2},
		RecType{"Bajingan",20,3.2},
		RecType{"Cuki",85,3.2},
		RecType{"Dick",1,3.2},
		RecType{"Entod",3,3.2},
		RecType{"Fuck",4,3.2},
		RecType{"Goblok",20,3.2},
	}

	fmt.Println(rmax(isinya))
	fmt.Println(imin(isinya))
	fmt.Println(found(isinya, "Ancok"))
	fmt.Println(pos(isinya, "Goblok"))
}

func rmax(data ArrType) int{
	var max int = data[0].f2
	var index = 0

	for i, val := range data{
		if max < val.f2 {
			max = val.f2
			index = i
		}
	}

	return index
}

func imin(data ArrType) int{
	var min int = data[0].f2
	var index = 0

	for i, val := range data{
		if min > val.f2 && val.f2 != 0{
			min = val.f2
			index = i
		}
	}

	return index
}

func found(data ArrType, key string) bool{
	for _, val := range data{
		if val.f1 == key{
			return true
		}
	}

	return false
}

func pos(data ArrType, key string) int {
	var leftIndex int = 0
	var rightIndex int = N
	var found bool = false

	for leftIndex < rightIndex  && !found{
		var median = (leftIndex+rightIndex) / 2
		
		if data[median].f1 < key {
			leftIndex = median + 1
		} else if data[median].f1 > key{
			rightIndex = median
		} else{
			found = true
			return median
		}
	}
	return -1
}