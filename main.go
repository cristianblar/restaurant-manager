package main

import (
	"time"
)

func main() {

	today := time.Now().Unix() // Default date

	todayData := fetchDayData(today)

}
