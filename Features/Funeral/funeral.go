package funeral

import (
	"fmt"
	"time"
)

var (
	year           string
	month          string
	day            string
	warnTime       string
	waitTime       time.Duration
	sendPeriodTime time.Duration
	incTime        time.Duration
)

func main() {
	warnTime = "02:00:00"
	waitTime = 15 * time.Minute
	sendPeriodTime = 1 * time.Minute
	incTime = 15 * time.Minute

	t := time.Now().Format("2006-01-02")
	t += " " + warnTime
	fmt.Println(t)
}
