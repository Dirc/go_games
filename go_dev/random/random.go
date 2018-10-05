package main

import (
	"fmt"
	"math/rand"
	"time"
)

var palet = [4]string{"red", "green", "blue", "yellow"}

func randomColor(palet [4]string) string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(palet))
	fmt.Println(len(palet), r, palet[r])
	return palet[r]
}

func main() {

	//palet := [4]string{"red", "green", "blue", "yellow"}

	fmt.Println(palet)
	//fmt.Println(rand.Intn(3))

	randomColor(palet)

	r := rand.Intn(len(palet))
	fmt.Println(len(palet), r, palet[r])

}
