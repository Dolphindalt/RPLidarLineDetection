package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	test()
}

func test() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("bababa")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var points []Point

	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")
		//quality, _ := strconv.Atoi(splitLine[0])
		angle, _ := strconv.ParseFloat(splitLine[1], 64)
		dist, _ := strconv.ParseFloat(splitLine[2], 64)
		PPoint := PolarPoint{dist, angle}
		points = append(points, CartesianPointFrom(PPoint))
	}

	arrayToImageSpace(points, 4000)

	for _, p := range points {
		fmt.Println(p.x, p.y)
	}

	h := NewHoughTransform(8000, 8000, points)

}
