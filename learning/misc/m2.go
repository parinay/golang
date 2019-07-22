package main

import (
	"time"
)

func two() {
	t := time.NewTicker(100)
	for range t.C {
		t.Stop()
	}
}
