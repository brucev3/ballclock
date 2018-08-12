package main

import (
	"fmt"
	"brucev.com/bclock/mods"
	"time"
)

func main() {
	start := time.Now()
	var balls, minutes = mods.Validateargs()
	var clock = mods.Clock{
		0,
		balls,
		minutes,
		mods.Minutes{Count:0,Min:0,Max:5,Balls:mods.NewStack(0)},
		mods.FiveMinutes{Count:0,Min:0,Max:12,Balls:mods.NewStack(0)},
		mods.Hours{Count:1,Min:1,Max:12,Balls:mods.NewStack(0)},
		mods.Bottom{Balls:mods.NewQueue(0)},
	}

	mods.Load(clock)
	mods.Run(clock)

	if clock.Minutes == 0 {
		//fmt.Println(clock.ElapsedDays)
		//fmt.Println(clock.Main.Balls.Q)
		//fmt.Println(clock.Hour.Balls.S)
		//fmt.Println(clock.FiveMin.Balls.S)
		//fmt.Println(clock.Min.Balls.S)
		//fmt.Println("")

		elapsed := time.Since(start)
		nanoseconds := float64(elapsed.Nanoseconds())
		milliseconds := nanoseconds/1000000
		seconds := milliseconds/1000

		//fmt.Println("elapsed is",elapsed)
		//fmt.Println("nanoseconds is",nanoseconds)
		//fmt.Println("milliseconds is",milliseconds)
		//fmt.Println("seconds is",seconds)

		fmt.Println(clock.Balls, "balls cycle after", clock.ElapsedDays, "days.")
		fmt.Println("Completed in", milliseconds, "milliseconds (",seconds,"seconds )")
	} else {
		state := mods.GetClockState(clock)

		//fmt.Println(clock.ElapsedDays)
		//fmt.Println(clock.Main.Balls.Q)
		//fmt.Println(clock.Hour.Balls.S)
		//fmt.Println(clock.FiveMin.Balls.S)
		//fmt.Println(clock.Min.Balls.S)
		//fmt.Println("")

		fmt.Println(string(state))
	}
}
