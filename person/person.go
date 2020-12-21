// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package person

import (
	"time"

	"bitbucket.com/paingha/data-structures/utils"
)

//Person - struct
type Person struct {
	FirstName  string
	LastName   string
	BirthDay   int
	BirthMonth int
	BirthYear  int
}

//BirthDate - struct method returns birth date of person
func (p Person) BirthDate() time.Time {
	return utils.ParseDate(p.BirthYear, p.BirthMonth, p.BirthDay, 0, 0, 0, 0, time.UTC)
}

//Age - method that returns the age of the Person
func (p Person) Age() float64 {
	t := p.BirthDate()
	currentTime := utils.SystemTime()
	age := utils.SubtractTime(currentTime, t)
	return utils.ConvertSecondsToYears(age)
}
