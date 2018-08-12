package main

import (
	"fmt"
	"brucev.com/bclock/mods"
	"encoding/json"
)

func main() {
	var balls, minutes = mods.Validateargs()
	var clock = mods.Clock{
		balls,
		minutes,
		mods.Minutes{Count:0,Min:0,Max:5,Balls:mods.NewStack(0)},
		mods.FiveMinutes{Count:0,Min:0,Max:12,Balls:mods.NewStack(0)},
		mods.Hours{Count:1,Min:1,Max:12,Balls:mods.NewStack(0)},
		mods.Bottom{Balls:mods.NewQueue(0)},
	}

	mods.Load(clock)
	var days, runtime = mods.Run(clock)

	if clock.Minutes == 0 {
		fmt.Println(clock.Balls, "balls cycle after", days, "days.")
		fmt.Println("Completed in", runtime, "milliseconds (y.yyy seconds)")
	} else {

		var clockstate = mods.Clockstate{[]int{},[]int{},[]int{},[]int{}}

		mainlength := len(clock.Main.Balls.Q)
		for i := 0; i < mainlength; i++ {
			ballnum := clock.Main.Balls.Q[i].Number
			clockstate.Main = append(clockstate.Main, ballnum)
		}

		hourlength := len(clock.Hour.Balls.S)
		for i := 0; i < hourlength; i++ {
			ballnum := clock.Hour.Balls.S[i].Number
			clockstate.Hour = append(clockstate.Hour, ballnum)
		}

		fiveminlength := len(clock.FiveMin.Balls.S)
		for i := 0; i < fiveminlength; i++ {
			ballnum := clock.FiveMin.Balls.S[i].Number
			clockstate.FiveMin = append(clockstate.FiveMin, ballnum)
		}

		minlength := len(clock.Min.Balls.S)
		for i := 0; i < minlength; i++ {
			ballnum := clock.Min.Balls.S[i].Number
			clockstate.Min = append(clockstate.Min, ballnum)
		}

		state, err := json.Marshal(clockstate)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(clock.Main.Balls.Q)
		fmt.Println(clock.Hour.Balls.S)
		fmt.Println(clock.FiveMin.Balls.S)
		fmt.Println(clock.Min.Balls.S)

		fmt.Println(string(state))
	}
}
