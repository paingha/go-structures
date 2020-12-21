// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

//Months - array of months for validation
var Months = map[int]string{
	1:  "Jan",
	2:  "Feb",
	3:  "Mar",
	4:  "Apr",
	5:  "May",
	6:  "Jun",
	7:  "Jul",
	8:  "Aug",
	9:  "Sept",
	10: "Oct",
	11: "Nov",
	12: "Dec",
}

//Days - array of days in a month for validation
var Days = map[int]int{
	1:  31,
	2:  29,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

//BirthDayValidator - validates the birth day input and makes sure it is correct
func BirthDayValidator(day, month int) bool {
	if value, exists := Days[month]; exists {
		if day <= value {
			return true
		}
	}
	return false
}

//BirthMonthValidator - validates the birth month input and makes sure it is correct
func BirthMonthValidator(month int) bool {
	if month <= 12 {
		if _, exists := Months[month]; exists {
			return true
		}
	}
	return false
}

//BirthYearValidator - validates the birth year input and makes sure it is correct
func BirthYearValidator(year int) bool {
	if year <= SystemTime().Year() {
		return true
	}
	return false
}
