/*
	Name	: Muhammad Ilham Mubarak
	SID		: 1301194276
	This program is used to calculate the closest distance from the given input points
	The input are terminated when its given pair of 0.0 points
	It will compare all given points
*/
package main

import "fmt"
import "math"

type Point struct{
	x float64
	y float64
}

const N = 10000

var points [N]Point
var numpoints int

func distance(p1, p2 Point) float64{
	return math.Sqrt(((p1.x - p2.x)*(p1.x - p2.x)) + ((p1.y - p2.y)*(p1.y - p2.y)))
}

func readPoints(){
	var point1, point2 float64
	numpoints = 0
	fmt.Scanln(&point1, &point2)

	for point1!=0 || point2!=0{
		points[numpoints] = Point{point1,point2}
		numpoints++
		fmt.Scanln(&point1, &point2)
	}
}

func getClosestPoint(p1, p2 *Point){
	var min = distance(points[0], points[1])

	for i := 0; i < numpoints; i++ {
		for j := i + 1; j < numpoints ; j++ {
			if distance(points[i], points[j]) < min {
				min = distance(points[i], points[j])
				*p1 = points[i]
				*p2 = points[j]
			}	
		}
	}
}

func main(){
	var p1, p2 Point
	readPoints()
	getClosestPoint(&p1, &p2)

	fmt.Println(p1, p2)

	fmt.Printf("Closest Point are (%.1f,%.1f) and (%.1f,%.1f) and the distance is %.1f\n", p1.x, p1.y, p2.x, p2.y, distance(p1, p2))
	fmt.Printf("Name \t: Muhammad Ilham Mubarak\n")
	fmt.Printf("SID \t: 1301194276\n")
}