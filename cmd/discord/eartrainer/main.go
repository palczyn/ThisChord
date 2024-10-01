package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var token string

type Interval int64

const (
	Unison Interval = iota
	MinorSecond
	MajorSecond
	MinorThird
	MajorThird
	PerfectFourth
	Tritone
	PerfectFifth
	MinorSixth
	MajorSixth
	MinorSeventh
	MajorSeventh
	Octave
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error while loading env file: %v", err)
		return
	}

	var ok bool
	token, ok = os.LookupEnv("API_TOKEN")

	if !ok || token == "" {
		fmt.Println("No token provided. Please ensure env is set.")
		return
	}
}
