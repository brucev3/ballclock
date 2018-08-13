package mods

import (
	"testing"
)

func TestIsInitialOrdering1(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var new1 = "{\"Min\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31],\"FiveMin\":[],\"Hour\":[],\"Main\":[]}"

	resnew1 := IsInitialOrdering(initial, new1)
		if resnew1 {
			//t.Error("Result is incorrect, got: %t, want: %t.", resnew, false)
			t.Error("Result is incorrect, got: true, want: false.")
		}
	}

func TestIsInitialOrdering2(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var new2 = "{\"Min\":[1,2,3],\"FiveMin\":[],\"Hour\":[],\"Main\":[4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"

	resnew2 := IsInitialOrdering(initial, new2)
		if resnew2 {
			t.Error("Result is incorrect, got: true, want: false.")
		}
	}

func TestIsInitialOrdering3(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var match = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"

	resmatch := IsInitialOrdering(initial, match)
		if !resmatch {
			t.Error("Result is incorrect, got: false, want: true.")
		}
	}