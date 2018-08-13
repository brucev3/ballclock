package mods

import (
	"errors"
	"encoding/json"
	"fmt"
	"os"
)

func Load(c Clock) {
	// load the balls
	ballnum := 1
	for ballnum <= c.Balls {
		//fmt.Println("new ballnum is:", ballnum)
		c.Main.Balls.Enqueue(NewBall(ballnum))
		ballnum += 1
	}
	//fmt.Println(c.Main.Balls) //TODO debug
}

func Run(c Clock) (days int) {

	minutes := 0
	days = 0

	if len(c.Main.Balls.Q) > 0 {
		initialstate :=  string(GetClockState(c))

		if c.Minutes == 0 {
			// repeatedly load balls from Main to Min until the first ball
			// becomes next up in the Main queue then return
			MinuteTick(c)
			minutes++

			nextnum := c.Main.Balls.Peek()
			for nextnum > 1 {
				MinuteTick(c)
				minutes++
				days = ComputeDays(minutes)
				nextnum = c.Main.Balls.Peek()
				//fmt.Println("Next loop ball:", nextnum)

				// bail out?
				if nextnum == 1 {
					currentstate := string(GetClockState(c))
					if IsInitialOrdering(initialstate, currentstate) {return}
				}
			}
		} else {
			// repeatedly load balls from Main to Min until c.Minutes expires then return
			MinuteTick(c)
			minutes++

			// bail out, minutes are expired
			if c.Minutes == minutes {return}

			nextnum := c.Main.Balls.Peek()
			for nextnum > 1 {
				MinuteTick(c)
				minutes++
				days = ComputeDays(minutes)
				nextnum = c.Main.Balls.Peek()
				//fmt.Println("Next loop ball:", nextnum)

				// bail out, minutes are expired
				if c.Minutes == minutes {return}
			}
		}
	} else {
		fmt.Println("Error: no balls installed. How did we get here???")
		os.Exit(1)
	}

	return days
}

func MinuteTick(c Clock) {

	ball, err := c.Main.Balls.Dequeue()
	if err == nil && ball.Number != -1 {
		c.Min.Balls.Push(ball)
	}
}

func GetClockState(c Clock) []byte {

		var clockstate = Clockstate{[]int{},[]int{},[]int{},[]int{}}

		mainlength := len(c.Main.Balls.Q)
		for i := 0; i < mainlength; i++ {
			ballnum := c.Main.Balls.Q[i].Number
			clockstate.Main = append(clockstate.Main, ballnum)
		}

		hourlength := len(c.Hour.Balls.S)
		for i := 0; i < hourlength; i++ {
			ballnum := c.Hour.Balls.S[i].Number
			clockstate.Hour = append(clockstate.Hour, ballnum)
		}

		fiveminlength := len(c.FiveMin.Balls.S)
		for i := 0; i < fiveminlength; i++ {
			ballnum := c.FiveMin.Balls.S[i].Number
			clockstate.FiveMin = append(clockstate.FiveMin, ballnum)
		}

		minlength := len(c.Min.Balls.S)
		for i := 0; i < minlength; i++ {
			ballnum := c.Min.Balls.S[i].Number
			clockstate.Min = append(clockstate.Min, ballnum)
		}

		state, err := json.Marshal(clockstate)
		if err != nil {
			fmt.Println(err)
		}

	return state
}

func (s *Stack) Push(b Ball) {
	s.S = append(s.S, b)
}

func (s *Stack) Pop() (Ball, error) {
	l := len(s.S)
	if l == 0 {
		return NewBall(-1), errors.New("empty stack")
	}
	result := s.S[l-1]
	s.S = s.S[:l-1]
	return result, nil
}

func (q *Queue) Enqueue(b Ball) {
	q.Q = append(q.Q, b)
}

func (q *Queue) Dequeue() (Ball, error) {
	if len(q.Q) == 0 {
		return NewBall(-1), errors.New("empty queue")
	}
	result := q.Q[0]
	q.Q = q.Q[1:]
	return result, nil
}

func (q *Queue) Peek() int {
	if len(q.Q) != 0 {
		return q.Q[0].Number
	} else {
		return -1
	}
}
