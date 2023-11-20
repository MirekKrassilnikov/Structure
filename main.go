package main

import "time"

type workers struct {

	ID             int
	Name           string
	BirthDate      time.Time
	HireDate       time.Time
	IsEmployed     bool
	Salary         float64
}
func () {
	worker1 := workers{
		ID:         1,
		Name:       "Alex",
		BirthDate:  parseDate("1990-01-15"),
		HireDate:   parseDate("2022-03-20"),
		IsEmployed: true,
		Salary:     50000.0,
	}
	worker2 := workers{
		ID:         2,
		Name:       "John",
		BirthDate:  parseDate("1993-02-11"),
		HireDate:   parseDate("2020-02-01"),
		IsEmployed: true,
		Salary:     40000.0
	}
	worker3 := workers{
		ID:         3,
		Name:       "Tom",
		BirthDate:  parseDate("1950-06-09"),
		HireDate:   parseDate("2002-08-30"),
		IsEmployed: false,
		Salary:     20000.0
	}

	worker4 := workers{
		ID:         4,
		Name:       "Sam",
		BirthDate:  parseDate("1978-09-13"),
		HireDate:   parseDate("2010-06-05"),
		IsEmployed: false,
		Salary:     25000.0
	}
	printworkersInfo(worker1)
	printworkersInfo(worker2)
	printworkersInfo(worker3)
	printworkersInfo(worker4)
}