package main

import "testing"

func Test_userGuessedIs(t *testing.T) {
	type args struct {
		numberToGuess int
		userGuess     int
	}
	tests := []struct {
		name string
		args args
		want guessResult
	}{
		{"User guess is too low", args{10, 9}, lower},
		{"User guess is too high", args{10, 11}, higher},
		{"User guess is equal", args{10, 10}, equal},
	}
	for _, tt := range tests {
		if got := userGuessedIs(tt.args.numberToGuess, tt.args.userGuess); got != tt.want {
			t.Errorf("%q. userGuessedIs(%d, %d) = %v, want %v", tt.name, tt.args.numberToGuess, tt.args.userGuess, got, tt.want)
		}
	}
}

func Test_guessInStack(t *testing.T) {
	tests := []struct {
		name  string
		stack []int
		guess int
		want  bool
	}{
		{"Number is in stack", []int{1, 2, 3}, 2, true},
		{"Number is not in stack", []int{1, 2, 3}, 5, false},
	}
	for _, tt := range tests {
		if got := guessInStack(tt.guess, tt.stack); got != tt.want {
			t.Errorf("%q. guessInStack(), %d = %v, want %v", tt.name, tt.guess, got, tt.want)
		}
	}
}

func Test_getRandomNumber(t *testing.T) {
	type args struct {
		bottomLimit int
		topLimit    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Number is between 1 and 10", args{1, 10}},
		{"Number is between 20 and 100", args{20, 100}},
	}

	for _, tt := range tests {
		got := getRandomNumber(tt.args.bottomLimit, tt.args.topLimit)
		if got < tt.args.bottomLimit || got > tt.args.topLimit {
			t.Errorf("getRandomNumber(), bottom %d, top %d, got %d", tt.args.bottomLimit, tt.args.topLimit, got)
		}
	}
}
