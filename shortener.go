package main

import (
	"math/rand"
	"strings"
)

const shortLinkSize int = 5

var hexMap = [16]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f",
}

func GenerateShort() string {
	var short strings.Builder
	for i := 0; i < shortLinkSize; i++ {
		short.WriteString(hexMap[rand.Intn(16)])
	}
	return short.String()
}
