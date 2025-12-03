package main

import (
	"fmt"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start *big.Int
	end   *big.Int
}

func GetInput() (ranges []Range) {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.TrimSpace(string(file))

	for line := range strings.SplitSeq(input, ",") {
		var r Range
		split_range := strings.Split(line, "-")
		start, _ := strconv.ParseInt(split_range[0], 10, 0)
		end, _ := strconv.ParseInt(split_range[1], 10, 0)

		r.start = big.NewInt(start)
		r.end = big.NewInt(end)

		ranges = append(ranges, r)
	}

	return ranges
}

func AllEqual(s string, i int) bool {
	l := len(s)

	if l%i != 0 {
		return false
	}
	nparts := l / i

	started := false
	pattern := ""
	for t := range slices.Chunk([]byte(s), nparts) {
		if !started {
			started = true
			pattern = string(t)
			continue
		}

		if string(t) != pattern {
			return false
		}
	}

	return true
}

func Star1() {
	total := big.NewInt(0)
	iter := big.NewInt(0)
	one := big.NewInt(1)
	for _, r := range GetInput() {
		iter.Set(r.start)
		for iter.Cmp(r.end) != 1 {
			s := iter.String()

			if AllEqual(s, 2) {
				total.Add(total, iter)
			}
			iter.Add(iter, one)
		}
	}

	fmt.Println(total.String())
}

func Star2() {
	total := big.NewInt(0)
	iter := big.NewInt(0)
	one := big.NewInt(1)
	for _, r := range GetInput() {
		iter.Set(r.start)
		for iter.Cmp(r.end) != 1 {
			s := iter.String()

			invalid := false
			for i := 2; i <= len(s); i++ {
				if AllEqual(s, i) {
					invalid = true
					break
				}
			}

			if invalid {
				total.Add(total, iter)
			}
			iter.Add(iter, one)
		}
	}

	fmt.Println(total.String())
}

func main() {
	Star1()
	Star2()
}
