/*
	Name : Muhammad Ilham Mubarak
	Kelas: IF-43-INT
	SID : 1301194276
	This program is made to implement the insertion sort and selection sort. 
	Sorting the given array in ascending and descending order
*/ 

package main

const MAXSIZE = 10

type RecType struct{
	count int
	size int
}

type ArrType [MAXSIZE]RecType

func isSort(tab *ArrType, nSize int){
	
	var i int = 1
	for i < nSize {
		temp := tab[i].count
		var j int = i - 1
		for j >= 0 && tab[j].count > temp{
			tab[j+1].count = tab[j].count
			j--
		}

		tab[j+1].count = temp
		i++
	}
}

func mSort(tab *ArrType, nSize int){
	var i int = 0

	for i < nSize{
		var max int = i
		var j int =  i + 1

		for j < nSize {
			if tab[max].size < tab[j].size {
				max = j
			}
			j++
		}

		temp := tab[i].size
		tab[i].size = tab[max].size
		tab[max].size = temp
		i++
	}
}

func isFound(tab ArrType, n, v int) bool{
	var left int = 0
	var right int = n
	var found bool = false

	for left < right && !found{
		var med = (left + right) / 2

		if tab[med].count > v {
			right = med
		} else if tab[med].count < v {
			left = med + 1
		} else{
			found = true
		}
	}

	return found
}