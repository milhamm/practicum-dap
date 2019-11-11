package main

import "fmt"


func main(){

	var arrTopping [9]string = [9]string{
		"Jamur",
		"Paprika",
		"Nanas",
		"Jagung",
		"Daging Asap",
		"Sosis Sapi",
		"Sosis Ayam",
		"Seafood",
		"Mozzarella",
	}

	fmt.Println(" _____    _     _   _ ___ _            ")
	fmt.Println("|_   _|__| |___| | | | _ (_)________ _ ")	
	fmt.Println("  | |/ -_) |___| |_| |  _/ |_ /_ / _` |")	
	fmt.Println("  |_|\\___|_|    \\___/|_| |_/__/__\\__,_|")	

	fmt.Printf("+----+-------------+--------+\n")
	fmt.Printf("| No |   Topping   | Harga  |\n")
	fmt.Printf("+----+-------------+--------+\n")
	for i := 0; i < 9; i++ {
		fmt.Printf("|%-4d|%-13s|%-8s|\n",i+1, arrTopping[i], "Rp60.000")
		fmt.Printf("+----+-------------+--------+\n")
	}
}