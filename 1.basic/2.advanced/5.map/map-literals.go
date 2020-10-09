package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//定义一个map的 结构体
type kubord struct {
	Last float64
	next float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {

	fmt.Println(m)

	m := map[string]kubord{
		"one": kubord{
			31, 22.00,
		},
		"two": {
			44, 45.00,
		},
	}

	fmt.Println(m)

}
