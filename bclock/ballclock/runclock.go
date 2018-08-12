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

		mainlength := len(clock.Main.Balls.Q)
		var mainballs = mods.Main{}
		for i := 0; i < mainlength; i++ {
			ballnum := clock.Main.Balls.Q[i].Number
			mainballs.Main = append(mainballs.Main, ballnum)
		}
		hourlength := len(clock.Hour.Balls.S)
		var hourballs = mods.Hour{}
		for i := 0; i < hourlength; i++ {
			ballnum := clock.Hour.Balls.S[i].Number
			hourballs.Hour = append(hourballs.Hour, ballnum)
		}
		fiveminlength := len(clock.FiveMin.Balls.S)
		var fiveminballs = mods.FiveMin{}
		for i := 0; i < fiveminlength; i++ {
			ballnum := clock.FiveMin.Balls.S[i].Number
			fiveminballs.FiveMin = append(fiveminballs.FiveMin, ballnum)
		}
		minlength := len(clock.Min.Balls.S)
		var minballs = mods.Min{}
		for i := 0; i < minlength; i++ {
			ballnum := clock.Min.Balls.S[i].Number
			minballs.Min = append(minballs.Min, ballnum)
		}


		//main, err := json.Marshal(mainballs)
		////main, err := json.MarshalIndent(clock.Main.Balls.Q,"Main:","") //Marshal(clock.Main.Balls)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//hour, err := json.Marshal(hourballs)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fivemin, err := json.Marshal(fiveminballs)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//min, err := json.Marshal(minballs)
		//if err != nil {
		//	fmt.Println(err)
		//}


		var clockstate = mods.Clockstate{minballs.Min,fiveminballs.FiveMin,hourballs.Hour,mainballs.Main}
		state, err := json.Marshal(clockstate)
		if err != nil {
			fmt.Println(err)
		}



		//main, err := json.Marshal(clock.Main.Balls)
		////main, err := json.MarshalIndent(clock.Main.Balls.Q,"Main:","") //Marshal(clock.Main.Balls)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//hour, err := json.Marshal(clock.Hour.Balls.S)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fivemin, err := json.Marshal(clock.FiveMin.Balls.S)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//min, err := json.Marshal(clock.Min.Balls.S)
		//if err != nil {
		//	fmt.Println(err)
		//}
		fmt.Println(string(main))
		fmt.Println(string(hour))
		fmt.Println(string(fivemin))
		fmt.Println(string(min))
		fmt.Println(clock.Main.Balls.Q)
		fmt.Println(clock.Hour.Balls.S)
		fmt.Println(clock.FiveMin.Balls.S)
		fmt.Println(clock.Min.Balls.S)

		fmt.Println(string(state))
	}



}
