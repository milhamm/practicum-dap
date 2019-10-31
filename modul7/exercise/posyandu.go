package main

import "fmt"

var N int

func main()  {
	var bMin, bMax float64
	fmt.Print("Enter the Number of babies: ")
	fmt.Scanln(&N)

	berat := make([]float64, N)
	i := 0
	for i < N {
		baby := 0.0
		fmt.Printf("Baby %d: ", i+1)
		fmt.Scanln(&baby)
		berat[i] = baby
		i++
	}

	calcMinMax(berat, &bMin, &bMax)
	fmt.Printf("Minimum weight: %f\n", bMin)
	fmt.Printf("Maximum weight: %f\n", bMax)
	fmt.Printf("Average weight: %f\n", average(berat))

	fmt.Printf("Name \t: Muhammad Ilham Mubarak\n")
	fmt.Printf("SID \t: \n")
}

func calcMinMax(weights []float64, bMin, bMax *float64) {
	i := 0
	max := weights[0]
	min := weights[0]
	for i < N{
		if weights[i]>max {
			max = weights[i]
		}
		if weights[i]<min {
			min = weights[i]
		}
		i++
	}
	*bMax = max
	*bMin = min
}

func average(weights []float64) float64 {
	avg := 0.0
	for _, val := range weights{
		avg += val
	}
	return avg / float64(N)
}