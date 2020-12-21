// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"math"
	"time"
)

//SystemTime - returns current system time
func SystemTime() time.Time {
	return time.Now()
}

//SubtractTime - returns the difference between two dates in nanoseconds
func SubtractTime(currentTime, arbitraryTime time.Time) time.Duration {
	return currentTime.Sub(arbitraryTime)
}

//ConvertTimeToHours - converts a give time with hours, minutes and seconds to just hours
func ConvertTimeToHours(clock time.Duration) float64 {
	return clock.Hours()
}

//ConvertSecondsToYears - converts seconds to years value
func ConvertSecondsToYears(seconds time.Duration) float64 {
	return math.Floor(ConvertTimeToHours(seconds) / 8760)
}

//ParseDate - parse date to time and date type
func ParseDate(year, month, day, hour, minute, second, nanosecond int, timezone *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, timezone)
}
