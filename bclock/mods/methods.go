package mods

import (
	"errors"
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

func Run(c Clock) (days, runtime float64) {



















	days = 0.1
	runtime = 0.1
	return days, runtime
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

func (q *Queue) Dequeue(b Ball) (Ball, error) {
	if len(q.Q) == 0 {
		return NewBall(-1), errors.New("empty queue")
	}
	result := q.Q[0]
	q.Q = q.Q[1:]
	return result, nil
}