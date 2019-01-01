package main

import (
	"fmt"
	"github.com/grsmv/goweek"
	"time"
)

func customizedFunction() {

	birthday := time.Date(2019, 4, 7, 23, 59, 59, 0, time.UTC)
	day := birthday.Day()
	fmt.Println(Whatday(day))

}

func goWeekPractice() {
	week, _ := goweek.NewWeek(2019, 13)
	fmt.Println(week.Number)
}
