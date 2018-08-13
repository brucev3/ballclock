package mods

import (
	"testing"
)

func TestIsInitialOrdering1(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var mismatch = "{\"Min\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31],\"FiveMin\":[],\"Hour\":[],\"Main\":[]}"

	result := IsInitialOrdering(initial, mismatch)
	if result {
		t.Errorf("Result is incorrect, got: %t, want: %t.", result, false)
	}
}

func TestIsInitialOrdering2(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var mismatch = "{\"Min\":[1,2,3],\"FiveMin\":[],\"Hour\":[],\"Main\":[4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"

	result := IsInitialOrdering(initial, mismatch)
	if result {
		t.Errorf("Result is incorrect, got: %t, want: %t.", result, false)
	}
}

func TestIsInitialOrdering3(t *testing.T) {
	var initial = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"
	var match = "{\"Min\":[],\"FiveMin\":[],\"Hour\":[],\"Main\":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31]}"

	result := IsInitialOrdering(initial, match)
	if !result {
		t.Errorf("Result is incorrect, got: %t, want: %t.", result, true)
	}
}

func TestComputeDays1(t *testing.T) {
	var minutes = 1430

	result := ComputeDays(minutes)
	if result != 0 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestComputeDays2(t *testing.T) {
	var minutes = 1440

	result := ComputeDays(minutes)
	if result != 1 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestComputeDays3(t *testing.T) {
	var minutes = 1445

	result := ComputeDays(minutes)
	if result != 1 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestComputeDays4(t *testing.T) {
	var minutes = 2870

	result := ComputeDays(minutes)
	if result != 1 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, 1)
	}
}

func TestComputeDays5(t *testing.T) {
	var minutes = 2885

	result := ComputeDays(minutes)
	if result != 2 {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, 2)
	}
}
