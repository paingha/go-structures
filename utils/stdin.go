// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//StdinString - function to get string input
func StdinString(prompt string, p *string) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		*p = UppercaseName(scanner.Text())
	}
}

//StdinInt - recursive function to get int input
func StdinInt(prompt string, p *int, validator func(int) bool) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		data := scanner.Text()
		i, err := strconv.Atoi(data)
		if err != nil || !validator(i) {
			fmt.Println("<< Invalid Input >>")
			StdinInt(prompt, p, validator)
		} else {
			*p = i
		}
	}
}

//StdinIntDay - recursive function to get int input
func StdinIntDay(prompt string, p *int, month int, validator func(int, int) bool) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		data := scanner.Text()
		i, err := strconv.Atoi(data)
		if err != nil || !validator(i, month) {
			fmt.Println("<< Invalid Input >>")
			StdinIntDay(prompt, p, month, validator)
		} else {
			*p = i
		}
	}
}
