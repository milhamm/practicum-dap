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
		RecType{"Aang",2,3.1},
		RecType{"Bambang",20,40.1},
		RecType{"Cecep",85,1.1},
		RecType{"Didit",1,1.2},
		RecType{"Erlang",3,5.2},
		RecType{"Farhan",4,7.2},
		RecType{"Gunawan",20,8.2},
	}

	fmt.Println(rmax(isinya))
	fmt.Println(imin(isinya))
	fmt.Println(found(isinya, "Aang"))
	fmt.Println(pos(isinya, "Erlang"))
}

func rmax(data ArrType) float64{
	var max float64 = data[0].f3

	for _, val := range data{
		if max < val.f3 {
			max = val.f3
		}
	}

	return max
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

	for leftIndex < rightIndex{
		var median = (leftIndex+rightIndex) / 2
		
		if data[median].f1 < key {
			leftIndex = median + 1
		} else if data[median].f1 > key{
			rightIndex = median
		} else{
			return median
		}
	}
	return -1
}