package challenge

import (
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {
	// TODO: your code here
	return -1.0
}

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
