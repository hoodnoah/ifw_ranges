package types

import (
	"time"
)

type TimeOfDay struct {
	Hour   uint
	Minute uint
}

type RangeDayHours struct {
	Date      time.Time
	StartTime TimeOfDay
	EndTime   TimeOfDay
}

type RangeHours struct {
	Range string
	Hours []RangeDayHours
}

type PartialRangeDayHours struct {
	Month     time.Month
	Day       int
	StartHour int
	StartMin  int
	EndHour   int
	EndMin    int
}

func (p *PartialRangeDayHours) ResolveYear(ref time.Time) RangeDayHours {
	year := ref.Year()

	// if the parsed month is January, and the reference is December, assume a new year
	if p.Month == time.January && ref.Month() == time.December {
		year += 1
	}

	loc := ref.Location()
	date := time.Date(year, p.Month, p.Day, 0, 0, 0, 0, loc)
	start := time.Date(year, p.Month, p.Day, p.StartHour, p.StartMin, 0, 0, loc)
	end := time.Date(year, p.Month, p.Day, p.EndHour, p.EndMin, 0, 0, loc)

	return RangeDayHours{
		Date:      date,
		StartTime: TimeOfDay{Hour: uint(start.Hour()), Minute: uint(start.Minute())},
		EndTime:   TimeOfDay{Hour: uint(end.Hour()), Minute: uint(end.Minute())},
	}
}
