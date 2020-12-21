// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"bitbucket.com/paingha/data-structures/person"
	"bitbucket.com/paingha/data-structures/utils"
)

func main() {
	me := person.Person{}
	utils.StdinString("What is your First Name: ", &me.FirstName)
	utils.StdinString("What is your Last Name: ", &me.LastName)
	utils.StdinInt("What is your Birth Month: ", &me.BirthMonth, utils.BirthMonthValidator)
	utils.StdinIntDay("What is your Birth Day: ", &me.BirthDay, me.BirthMonth, utils.BirthDayValidator)
	utils.StdinInt("What is your Birth Year: ", &me.BirthYear, utils.BirthYearValidator)
	fmt.Printf("%v %v, age %v was born on %v-%v-%v and today is %v-%v-%v \n",
		me.FirstName, me.LastName, me.Age(), me.BirthMonth, me.BirthDay,
		me.BirthYear, utils.SystemTime().Month(), utils.SystemTime().Day(),
		utils.SystemTime().Year())
}
