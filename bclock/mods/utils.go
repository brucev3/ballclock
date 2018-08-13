package mods

import (
	"fmt"
	"flag"
	"os"
	"strconv"
)

func Validateargs () (balls, minutes int) {

	minutes = 0
	flag.Parse()

	alength := len(flag.Args())
	switch {
	case alength == 0:
		fmt.Println("Error: missing input","\nPlease provide valid number of balls, 27-127\nlike:\n   $ ballclock 30")
		os.Exit(1)
	case alength == 1:
		i, err := strconv.Atoi(flag.Arg(0))
		if err == nil {
			balls = i
		} else {
			fmt.Println("Error: ",err, "\nPlease provide valid number of balls, 27-127\nlike:\n   $ ballclock 30")
			os.Exit(1)
		}
	default:
		i, erri := strconv.Atoi(flag.Arg(0))
		if erri == nil {
			balls = i
		} else {
			fmt.Println("Error: ",erri, "\nPlease provide valid number of balls, 27-127\nlike:\n   $ ballclock 30")
			os.Exit(1)
		}
		j, errj := strconv.Atoi(flag.Arg(1))
		if errj == nil {
			minutes = j
		} else {
			fmt.Println("Error: ",errj, "\nPlease provide valid number of minutes\nlike:\n   $ ballclock 30 20")
			os.Exit(1)
		}
	}
	// Check balls range
	switch {
	case balls < 27:
		fmt.Println("Error: not enough balls:",balls,"\nPlease provide valid number of balls, 27-127\nlike:\n   $ ballclock 30")
		os.Exit(1)
	case balls > 127:
		fmt.Println("Error: too many balls:",balls,"\nPlease provide valid number of balls, 27-127\nlike:\n   $ ballclock 30")
		os.Exit(1)
	}
	return balls, minutes
}

func ComputeDays(minutes int) (days int) {

	return minutes/1440
}

func IsInitialOrdering(initial, current string) bool {

	if current == initial {
		return true
	}
	return false
}

func NewBall(number int) Ball {
	return Ball{number}
}

func NewStack(size int) *Stack {
	return &Stack {make([]Ball, size)}
}

func NewQueue(size int) *Queue {
	return &Queue {make([]Ball, size)}
}
