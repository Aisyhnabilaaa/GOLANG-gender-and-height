package main

import "fmt"
func main () {
	var gender string
	var height int
	var result float64
		fmt.Println("Masukkan gender : ")
		fmt.Scan(&gender)
		fmt.Println("Masukkan height : ")
		fmt.Scan(&height)

		if gender == "laki laki"{
			result = (float64(height)-100) - ((float64(height)-100) * 0.1)
			fmt.Println(result)
		}
		 if gender == "perempuan"{
			result = (float64(height)-100) - ((float64(height)-100) * 0.15)
			fmt.Println(result)
		 }
}