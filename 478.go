package main

import (
	"math"
	"math/rand"
)

// Given the radius and x-y positions of the center of a circle, write a function randPoint which generates a uniform random point in the circle.
//
// Note:
//
//     input and output values are in floating-point.
//     radius and x-y position of the center of the circle is passed into the class constructor.
//     a point on the circumference of the circle is considered to be in the circle.
//     randPoint returns a size 2 array containing x-position and y-position of the random point, in that order.
//
// Example 1:
//
// Input:
// ["Solution","randPoint","randPoint","randPoint"]
// [[1,0,0],[],[],[]]
// Output: [null,[-0.72939,-0.65505],[-0.78502,-0.28626],[-0.83119,-0.19803]]
//
// Example 2:
//
// Input:
// ["Solution","randPoint","randPoint","randPoint"]
// [[10,5,-7.5],[],[],[]]
// Output: [null,[11.52438,-8.33273],[2.46992,-16.21705],[11.13430,-12.42337]]
//
// Explanation of Input Syntax:
//
// The input is two lists: the subroutines called and their arguments. Solution's constructor has three arguments, the radius, x-position of the center, and y-position of the center of the circle. randPoint has no arguments. Arguments are always wrapped with a list, even if there aren't any.

type Solution struct {
	Radius float64
	X, Y   float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		Radius: radius,
		X:      x_center,
		Y:      y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	x := this.X - this.Radius
	y := this.Y - this.Radius

	for true {
		newX := x + rand.Float64()*this.Radius*2
		newY := y + rand.Float64()*this.Radius*2

		distX := newX - this.X
		distY := newY - this.Y

		if math.Sqrt(distX*distX+distY*distY) <= this.Radius {
			return []float64{newX, newY}
		}
	}

	return []float64{}
}

type Solution1 struct {
	Radius float64
	X, Y   float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution1 {
	return Solution1{
		Radius: radius,
		X:      x_center,
		Y:      y_center,
	}
}

func (this *Solution1) RandPoint() []float64 {
	dist := this.Radius * rand.Float64()
	angle := rand.Float64() * 360

	if angle <= 90 {
		x := this.X + math.Cos(dist)
		y := this.Y + math.Sin(dist)
		return []float64{x, y}
	} else if angle <= 180 {
		x := this.X - math.Cos(dist)
		y := this.Y + math.Sin(dist)
		return []float64{x, y}
	} else if angle <= 270 {
		x := this.X - math.Cos(dist)
		y := this.Y - math.Sin(dist)
		return []float64{x, y}
	}

	x := this.X + math.Cos(dist)
	y := this.Y - math.Sin(dist)

	return []float64{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

//	Notes
//	1.	inspired form solution, map square area to circle area
