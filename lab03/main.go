package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main(){
	var n int
	var counter int

	println("How many squares generate?")
	fmt.Scanf("%d", &n)

	for index := 0; index < n; index++ {
		square := Square{
			a: Point{rand.Intn(10), rand.Intn(10)},
			b: Point{rand.Intn(10), rand.Intn(10)},
			c: Point{rand.Intn(10), rand.Intn(10)},
			d: Point{ rand.Intn(10), rand.Intn(10)},
		}

		println("Square #" + strconv.Itoa(index))
		square.show()
		counter += checkIfSquare(square)
	}
	println("Generated " + strconv.Itoa(counter) + " correct squares.")
}

func checkIfSquare(square Square) int {
	var ab = math.Sqrt(math.Pow(float64(square.a.coordinateX - square.b.coordinateX), 2) + math.Pow(float64(square.a.coordinateY - square.b.coordinateY),2))
	var bc = math.Sqrt(math.Pow(float64(square.b.coordinateX - square.c.coordinateX), 2) + math.Pow(float64(square.b.coordinateY - square.c.coordinateY),2))
	var ad = math.Sqrt(math.Pow(float64(square.a.coordinateX - square.d.coordinateX), 2) + math.Pow(float64(square.a.coordinateY - square.d.coordinateY),2))
	var cd = math.Sqrt(math.Pow(float64(square.c.coordinateX - square.d.coordinateX), 2) + math.Pow(float64(square.c.coordinateY - square.d.coordinateY),2))

	if ab == bc && ab == ad && ab == cd && bc == ad && bc == cd && ad == cd {
		println("This square is geometrically correct.")
		return 1
	} else {
		return 0
	}

	return 0
}
