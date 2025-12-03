package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


const (
	DirectionLeft  = -1
	DirectionRight = 1
)

type Rotation struct {
	direction int
	distance  int
	str string
}

func GetInput() (rots []Rotation) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		var r Rotation
		line := scanner.Text()

		r.direction = DirectionRight
		if line[0] == 'L' {
			r.direction = DirectionLeft
		}
		
		n, _ := strconv.Atoi(line[1:])
		r.distance = n
		r.str = line
		rots = append(rots, r)
	}

	return rots
}

func Star1(){
	rots := GetInput()

	dial := 50
	count := 0
	for _, r := range rots {
		dial = (dial+r.direction*r.distance)%100

		if dial == 0 {
			count++
		}
	}

	fmt.Println(count)
}


func Star2(){
	rots := GetInput()

	dial := 50
	count := 0
	for _, r := range rots {
		count += r.distance/100
		new_dial := dial+r.direction*(r.distance%100)
		if (new_dial <= 0 || new_dial >= 100) && dial != 0 {
			count++	
		}
		dial = (new_dial % 100 + 100) % 100
	}

	fmt.Println(count)
}

func main() {
	Star1()
	Star2()
}
