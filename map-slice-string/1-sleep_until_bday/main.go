package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	p, err:= time.Parse(expectedFormat,target)
	if err != nil || time.Now().After(p) {
     log.Fatalf("invalid date")
	}
	return p;
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	return float64(time.Until(target).Hours()/24)
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}


// $ go run main.go -bday 2023-11-01
// 2023/10/01 15:40:32 You have 30 sleeps until your birthday. Hurray!