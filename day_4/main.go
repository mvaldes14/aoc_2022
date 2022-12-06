package main

import (
	U "aoc2022"
	"fmt"
)

func main() {
	data := U.ReadFile()
  numberOfContained := 0
  numberOfOverlaps := 0
	for _, line := range data {
    var startFirst, endFirst, startSecond, endSecond int
    fmt.Sscanf(line, "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)
		if startSecond >= startFirst && endSecond <= endFirst || startFirst >= startSecond && endFirst <= endSecond{
			numberOfContained++
		}
    if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond{
			numberOfOverlaps++
		}
	}
  fmt.Println(numberOfContained)
  fmt.Println(numberOfOverlaps)
}
