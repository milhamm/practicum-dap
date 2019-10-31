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


	if totalTime1 == totalTime2 && totalProb1 == totalProb2 {
		fmt.Printf("Both are co winners")
	} else{
		if totalProb1>totalProb2 {
			printWinner(name1, totalProb1, totalTime1)
		} else if totalProb2>totalProb1 {
			printWinner(name2, totalProb2, totalTime2)
		} else{
			if disq(totalTime1){
				fmt.Printf("%s is disqualified\n", name1)
			} else{
				printWinner(name1, totalProb1, totalTime1)
			}

			if disq(totalTime2){
				fmt.Printf("%s is disqualified\n", name2)
			} else{
				printWinner(name2, totalProb2, totalTime2)
			}
		}
	}
}

func calculateScore(prb, score *int){
	var time int
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&time)
		if isValid(time) {
			*score += time
			*prb += 1
		} else{
			fmt.Printf("Time is not Valid\n")
		}
	}
}

func printWinner(name string, probSolved, score int){
	fmt.Printf("%s %d %d\n", name, probSolved, score)
}

func disq(totalTime int)bool{
	return totalTime>301
}

func isValid(time int) bool{
	return time>0 && time<301
}