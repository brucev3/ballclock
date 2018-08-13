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
		balls,
		minutes,
		mods.Minutes{Count:0,Min:0,Max:5,Balls:mods.NewStack(0)},
		mods.FiveMinutes{Count:0,Min:0,Max:12,Balls:mods.NewStack(0)},
		mods.Hours{Count:1,Min:1,Max:12,Balls:mods.NewStack(0)},
		mods.Bottom{Balls:mods.NewQueue(0)},
	}

	mods.Load(clock)
	//mods.Load(clock)
	var days = mods.Run(clock)

	if clock.Minutes == 0 {

		elapsed := time.Since(start)
		nanoseconds := float64(elapsed.Nanoseconds())
		milliseconds := nanoseconds/1000000
		seconds := milliseconds/1000

		fmt.Println(clock.Balls, "balls cycle after", days, "days.")
		fmt.Println("Completed in", milliseconds, "milliseconds (",seconds,"seconds )")

		// TODO temporary for debug only
		fmt.Println(string(mods.GetClockState(clock)))

	} else {

		state := mods.GetClockState(clock)

		elapsed := time.Since(start)
		nanoseconds := float64(elapsed.Nanoseconds())
		milliseconds := nanoseconds/1000000
		seconds := milliseconds/1000

		fmt.Println(string(state))
		fmt.Println("Completed in", milliseconds, "milliseconds (",seconds,"seconds )")
	}
}
