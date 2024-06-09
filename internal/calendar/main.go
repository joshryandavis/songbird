package calendar

import (
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/gb"
	"time"
)

func NewCalendar() *cal.BusinessCalendar {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(gb.Holidays...)
	return c
}

func GetBankWorkingDay(c *cal.BusinessCalendar, date time.Time) time.Time {
	holiday, _, _ := c.IsHoliday(date)
	weekend := !c.IsWorkday(date)
	for holiday || weekend {
		date = date.AddDate(0, 0, 1)
		holiday, _, _ = c.IsHoliday(date)
		weekend = !c.IsWorkday(date)
	}
	return date
}
