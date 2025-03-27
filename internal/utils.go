package internal

import (
	"fmt"
	"time"
)

func PrintlnWithElapsedTime(string string, startTime time.Time) {
	elapsedTime := time.Since(startTime).Milliseconds()
	fmt.Println(fmt.Sprintf("%s (%dms)", string, elapsedTime))
}
