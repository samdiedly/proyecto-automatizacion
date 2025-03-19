package main

import (
	"fmt"
	"time"
)

func main() {
	startDate := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	var dates []string

	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("Monday, 02 January 2006"))
	}

	for _, date := range dates {
		fmt.Println(date)
	}
}
