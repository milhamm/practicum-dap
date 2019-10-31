package main

import "fmt"

func main(){
	var name1, name2 string
	var totalProb1, totalProb2 int
	var totalTime1, totalTime2 int

	fmt.Scan(&name1)
	calculateScore(&totalProb1, &totalTime1)
	fmt.Scan(&name2)
	calculateScore(&totalProb2, &totalTime2)

	if totalProb1>totalProb2 {
		printWinner(name1, totalProb1, totalTime1)
	} else if totalProb2>totalProb1 {
		printWinner(name2, totalProb2, totalTime2)
	}else{
		if totalTime1<totalTime2{
			printWinner(name1, totalProb1, totalTime1)
		} else {
			printWinner(name2, totalProb2, totalTime2)
		}
	}
}

func calculateScore(prb, score *int){
	var time int
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&time)
		if time < 301 {
			*score += time
			*prb += 1
		}
	}
}

func printWinner(name string, probSolved, score int){
	fmt.Printf("%s %d %d", name, probSolved, score)
}