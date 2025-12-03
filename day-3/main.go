package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() (banks [][]int) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var bank []int

		for _, c := range scanner.Text() {
			bank = append(bank, int(c-'0'))
		}
		banks = append(banks, bank)
	}

	return banks
}

func GetMaxBetween(bank []int, after, before int) (maxi, maxn int) {
	maxi = after
	maxn = bank[after]
	for i := after + 1; i < before; i++ {
		if bank[i] > maxn {
			maxi = i
			maxn = bank[i]
		}
	}

	return maxi, maxn
}

func Star1() {
	total := 0
	for _, bank := range GetInput() {
		i, tens := GetMaxBetween(bank, 0, len(bank)-1)
		_, ones := GetMaxBetween(bank, i+1, len(bank))
		total += tens*10 + ones
	}
	fmt.Println(total)
}

func Star2() {
	total := 0
	for _, bank := range GetInput() {
		value := 0
		start := -1
		for place := 11; place >= 0; place-- {
			i, digit := GetMaxBetween(bank, start+1, len(bank)-place)
			value = value*10 + digit
			start = i
		}
		total += value
	}
	fmt.Println(total)
}

func main() {
	Star1()
	Star2()
}
