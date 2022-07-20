//封装
package main

import (
	"fmt"
	"log"
	"errors"
	"unicode/utf8"
)

type Date struct {
	Year int
	Month int
	Day int
}

type Event struct {
	Title string
	Date
}

func main() {
	date := Date{
		Year:2022,
		Month:7,
		Day:15,
	}
	fmt.Println(date)

	d := Date{}
	err := d.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = d.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = d.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)


    fmt.Println(d.year(), d.month(), d.day())

	event := Event{}
	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	
	err = event.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetDay(23)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("asdgasdfgasg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
	fmt.Println(event.year())
	
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func (d *Date) year() int {
	return d.Year
}

func (d *Date) month() int {
	return d.Month
}

func (d *Date) day() int {
	return d.Day
}

func (e *Event) title() string {
	return e.Title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invild title")
	}
	e.Title = title
	return nil
}

