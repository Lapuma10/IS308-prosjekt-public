package rbMQ

import (
	"time"
	//"log"
)
/*
* Arguably doesnt belong in this package. 
*/
func IsToday(compDate string) bool {
	currentTime := time.Now()
	todayDate := currentTime.UTC().Format("2006-01-02")

	return compDate == todayDate
}

func IsXDaysAway(compDate string, days int) bool {
	currentTime := time.Now()
	targetTime := currentTime.Add(-time.Hour * 24 * time.Duration(days))

	layout := "2006-01-02"
	actualTime, _ := time.Parse(layout, compDate)

	//log.Printf("Current: %s, Target: %s, Actual: %s", currentTime.String(), targetTime.String(), actualTime.String())

	return actualTime == targetTime || actualTime.Before(targetTime)
}