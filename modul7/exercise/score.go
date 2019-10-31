package main

import "fmt"

func main(){
	var clubA, clubB string
	var winnerClub []string
	
	fmt.Print("Club A: ")
	fmt.Scan(&clubA)

	fmt.Print("Club B: ")
	fmt.Scan(&clubB)

	readMatch(&winnerClub, clubA, clubB)
	printResult(winnerClub)
}


func readMatch(winnerClub *[]string, clubA, clubB string)  {
	var scoreA, scoreB int
	var i int = 1
	
	fmt.Printf("Match %d: ", i)
	fmt.Scan(&scoreA, &scoreB)

	for scoreA >= 0 && scoreB >= 0 {
		if scoreA > scoreB {
			fillWinner(winnerClub, clubA)
		} else if scoreA < scoreB {
			fillWinner(winnerClub, clubB)
		} else {
			fillWinner(winnerClub, "Draw")
		}
		i++

		fmt.Printf("Match %d: ", i)
		fmt.Scan(&scoreA, &scoreB)
	}	
}

func fillWinner(winnerClub *[]string, name string){
	*winnerClub = append(*winnerClub, name)
}

func printResult(winnerClub []string){
	for i, val := range winnerClub {
		fmt.Printf("Result %d: %s\n", i+1, val)
	}

	fmt.Printf("Game has Ended")
}