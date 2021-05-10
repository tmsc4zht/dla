package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math"
	"math/rand"
)

func randInit() {
	var seed int64
	err := binary.Read(cryptorand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}
	rand.Seed(seed)
}

func randCircle(cr float64, cx, cy int) (int, int) {
	theta := rand.Float64() * math.Pi * 2
	x, y := cr*math.Cos(theta), cr*math.Sin(theta)
	return cx + int(x), cy + int(y)
}

func distance(cx, cy, x, y int) float64 {
	dx := float64(cx - x)
	dy := float64(cy - y)
	return math.Sqrt(dx*dx + dy*dy)
}
