/*
	Name	: Muhammad Ilham Mubarak
	Class	: IF-43-INT
	SID		: 1301194276

	DAP Final Project - Simple Sackson Game
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)


type pizza struct{

}

func main(){
	welcomeMessage()
	printScoreBoard()
}

func rollDice(dice *[4]int){
	for i := 0; i < 4;  i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		(*dice)[i] = rand.Intn(6) + 1
	}
}

func welcomeMessage(){
	var playerName string

	fmt.Printf("Welcome to Simple Simple Sexson Game :D\n")
	fmt.Printf("Please enter your name: ")
	fmt.Scan(&playerName)
}

func printScoreBoard(){
	fmt.Printf("+----------+------+\n")
	fmt.Printf("| Sum Pair | Mark |\n")
	fmt.Printf("+----------+------+\n")
	var i int = 0
	for i < 11 {
		fmt.Printf("|%-10d|%-6d|\n",i + 2, 0)
		fmt.Printf("+----------+------+\n")
		i++
	}
}

