package mods

type Ball struct {
	Number int
}

type Minutes struct {
	Count int
	Min int
	Max int
	Balls *Stack
}

type FiveMinutes struct {
	Count int
	Min int
	Max int
	Balls *Stack
}

type Hours struct {
	Count int
	Min int
	Max int
	Balls *Stack
}

type Bottom struct {
	Balls *Queue
}

type Clock struct {
	Balls int
	Minutes int
	Min Minutes
	FiveMin FiveMinutes
	Hour Hours
	Main Bottom
}

type Stack struct {
	S []Ball
}

type Queue struct {
	Q []Ball
}

// structs for formatting 'Mode 2 (Clock State)'
type Min struct {
	Min []int
}

type FiveMin struct {
	FiveMin []int
}

type Hour struct {
	Hour []int
}

type Main struct {
	Main []int
}

type Clockstate struct {
	Min []int
	FiveMin []int
	Hour []int
	Main []int
}