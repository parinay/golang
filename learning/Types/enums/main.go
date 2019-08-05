package main

import "fmt"

type Day int

const (
	MONDAY Day = 1 + iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func tellLikeItIs(day Day) {
	switch day {
	case MONDAY:
		fmt.Println("Mondays are drag")

	case TUESDAY, WEDNESDAY, THURSDAY:
		fmt.Println("Weekdays are so so")
	}
}
func main() {
	firstDay := MONDAY

	tellLikeItIs(firstDay)
}
