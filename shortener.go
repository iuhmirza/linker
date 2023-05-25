package main

import (
	"math/rand"
	"strings"
)

const SHORT_LINK_SIZE int = 5

var hexMap = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
}

func GenerateShort() string {
	var short strings.Builder
	for i := 0; i < SHORT_LINK_SIZE; i++ {
		short.WriteString(hexMap[rand.Intn(15)])
	}
	return short.String()
}
