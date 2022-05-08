package main

import (
	"testing"
)

func TestFizzBuzz_1(t *testing.T) {
	got := fizzbuzz(1)
	want := "1"

	if got != want {
		t.Errorf("fizzbuzz(1) \n got: %v \n want: %v\n", got, want)
	}
}

func TestFizzBuzz_3(t *testing.T) {
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3) \n got: %v \n want: %v\n", got, want)
	}
}

func TestFizzBuzz_5(t *testing.T) {
	got := fizzbuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("fizzbuzz(5) \n got: %v \n want: %v\n", got, want)
	}
}

func TestFizzBuzz_15(t *testing.T) {
	got := fizzbuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("fizzbuzz(15) \n got: %v \n want: %v\n", got, want)
	}
}
