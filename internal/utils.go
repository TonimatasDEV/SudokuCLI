package internal

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func printlnWithElapsedTime(string string, startTime time.Time) {
	elapsedTime := time.Since(startTime).Milliseconds()
	fmt.Println(fmt.Sprintf("%s (%dms)", string, elapsedTime))
}

func randomNum(limit int64) int8 {
	n, _ := rand.Int(rand.Reader, big.NewInt(limit))
	return int8(n.Int64())
}
