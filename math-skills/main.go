package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong size arguments")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	if len(data) < 1 {
		fmt.Println("Nothing to solve")
		return
	}

	datastr := strings.Split(string(data), "\n")

	if datastr[len(datastr)-1] == "" {
		datastr = datastr[:len(datastr)-1]
	}

	dataslice := []int{}
	for _, val := range datastr {
		if val != "" {
			valint, err := strconv.Atoi(val)
			if err != nil && val != "" {
				fmt.Println("Wrong symbols in data.txt")
				return
			}
			dataslice = append(dataslice, valint)
		} else {
			continue
		}

	}

	fmt.Println("Average:", math.Round(Average(dataslice)))
	fmt.Println("Median:", math.Round(Median(dataslice)))
	// fmt.Println("Variance", int(math.Round(Variance(dataslice))))
	fmt.Printf("Variance: %.0f\n", (Variance(dataslice)))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(math.Sqrt(Variance(dataslice))))

}

func Average(dataslice []int) float64 {
	var sum float64
	for _, num := range dataslice {
		sum += float64(num)
	}
	return sum / float64(len(dataslice))
}

func Median(dataslice []int) float64 {
	sort.Ints(dataslice)
	var mid float64
	if len(dataslice)%2 == 0 {

		mid = float64(dataslice[len(dataslice)/2-1]+dataslice[len(dataslice)/2]) / 2

		return mid

	}
	return float64(dataslice[len(dataslice)/2])
}

func Variance(dataslice []int) float64 {
	average := Average(dataslice)
	var sum float64
	for _, num := range dataslice {
		diff := float64(num) - average
		sum += diff * diff
	}
	variance := sum / float64(len(dataslice))
	return variance
}
