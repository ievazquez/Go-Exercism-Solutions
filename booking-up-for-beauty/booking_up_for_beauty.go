package booking

import (
    "time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    layouts := [3]string{"1/2/2006 15:04:05", "January 2, 2006 15:04:05", "Monday, January 2, 2006 15:04:05"}

    var t time.Time
    var err error 
	for _, layout:= range layouts {
        t, err = time.Parse(layout, date)
        if err == nil {
                break
        }
    }
    return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    today, myDate := time.Now(), Schedule(date)
    return today.After(myDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	myDate := Schedule(date)
    return myDate.Hour() >= 12 &&   myDate.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    dt := Schedule(date)
    return dt.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
 }

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15,  0, 0, 0, 0, time.UTC)
}
