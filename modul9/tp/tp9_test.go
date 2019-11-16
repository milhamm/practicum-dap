package main

import "testing"

// Testing for TP Module 9
var tab ArrType = ArrType{{5, 4},{6, 1},{10, 2},{1, 6},{4, 7},{8, 5},{2, 3},{7, 10},{3, 8},{9, 9}}
var expectedOutput ArrType = ArrType{{1,10},{2,9},{3,8},{4,7},{5,6},{6,5},{7,4},{8,3},{9,2},{10,1}}

var find int = 7
var expectedOutputFind = true

var findSecond int = 11
var expectedOutputFindSecond = false


func TestSorting(t *testing.T){

	isSort(&tab, 10)
	t.Logf("Sorted Ascending Array: %v", tab)
	mSort(&tab, 10)
	t.Logf("Sorted Descending Array: %v", tab)

	if tab != expectedOutput {
		t.Errorf("Failed! Expected value : %v", expectedOutput)
	}
}

func TestSearch(t *testing.T){
	t.Logf("Is 7 Found? %t", isFound(tab, 10, find))

	if isFound(tab, 10, find) != expectedOutputFind {
		t.Errorf("Failed! Expected value : %v", expectedOutputFind)
	}

	t.Logf("Is 11 Found? %t", isFound(tab, 10, findSecond))

	if isFound(tab, 10, findSecond) != expectedOutputFindSecond {
		t.Errorf("Failed! Expected value : %v", expectedOutputFindSecond)
	}
}

func BenchmarkInsertionSorting(b *testing.B) {
    for i := 0; i < b.N; i++ {
        isSort(&tab, 10)
    }
}

func BenchmarkSelectionSorting(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mSort(&tab, 10)
    }
}