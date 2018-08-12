package main

import (
	"fmt"
	"brucev.com/bclock/mods"
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
		fmt.Println(clock.Main.Balls.Q)
		fmt.Println(clock.Hour.Balls.S)
		fmt.Println(clock.FiveMin.Balls.S)
		fmt.Println(clock.Min.Balls.S)
		fmt.Println("")

		fmt.Println(clock.Balls, "balls cycle after", days, "days.")
		fmt.Println("Completed in", runtime, "milliseconds (y.yyy seconds)")
	} else {
		state := mods.GetClockState(clock)

		fmt.Println(clock.Main.Balls.Q)
		fmt.Println(clock.Hour.Balls.S)
		fmt.Println(clock.FiveMin.Balls.S)
		fmt.Println(clock.Min.Balls.S)
		fmt.Println("")

		fmt.Println(string(state))
	}
}
