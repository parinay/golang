package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	//m = make(map[string]Vertex)
	//m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
}
