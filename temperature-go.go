package main

import "fmt"

const ebulitionK float64 = 373.2

func main() {
	var tempK float64 = ebulitionK

	var tempC float64 = tempK - 273

	fmt.Println("The water ebulition temperature in °K = ", tempK)
	fmt.Println("The water ebulition temperature in °C = ", tempC)
}