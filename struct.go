package main

import (
	"fmt"
)

//类型通常在包级别定义
type part struct {
	description string
	count int
}

type car struct {
	name string 
	topSpeed float64
}

type subscriber struct {
	name string
	rate float64
	active bool
	HomeAddress Address
}

type Employee struct {
	Name string
	Salary float64
	*Address
}

type Address struct {
	Street string
	City string
	State string
	PostalCode string
}

type Liters float64
type Milliliters float64
type Gallons float64

func main() {
	var mark struct {
		name string
		age int
		sun bool
	}
	fmt.Printf("%#v\n", mark)

	mark.name = "刘诗琪"
	mark.age = 21
	mark.sun = true
	fmt.Printf("NAME: %v\n", mark.name)
	fmt.Printf("AGE: %v\n", mark.age)
	fmt.Printf("SUN: %v\n", mark.sun)


	var porsche car 
	porsche.name = "911"
	porsche.topSpeed = 323
	fmt.Printf("Porsche Name: %v\n", porsche.name)
	fmt.Printf("Porsche TopSpeed: %v\n", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex"
	bolts.count = 24
	showInfo(bolts)

	p := returnInfo("Hello World")
	fmt.Println(p.description, p.count)
	//fmt.Println("------------------------------")

	subscriber1 := madeInfo("Mark")
	applyDiscount(subscriber1)//这里不需要添加取址运算符，因为参数已经是一个struct指针
	printInfo(subscriber1)

	subscriber2 := madeInfo("Alex")
	printInfo(subscriber2)


	address1 := madeAdress("Oak")
	printAddress(address1)
	employee1 := madeEmployee("Tony")
	employee1.Address = address1
	printEmployee(employee1)


	//自定义具有底层基础类型的定义类型
	carFuel := Gallons(10.0)
	busFuel := Milliliters(20)
	fmt.Println(carFuel.ToLiters(), busFuel.ToLiters())

	soda := Liters(10)
    water := Milliliters(20)
	fmt.Println(soda.ToGallons(),water.ToGallons())





}


//自定义类型作为参数
func showInfo(p part) {
	fmt.Printf("Description: %v\n", p.description)
	fmt.Printf("Count: %v\n", p.count)
}

//自定义类型作为返回值
func returnInfo(info string) part {
	var p part
	p.description = info
	p.count = 100
	return p
}


func printInfo(s *subscriber) {
	fmt.Printf("**********%v**********\n", s.name)
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("Active: ", s.active)
}

func madeInfo(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func madeEmployee(name string) *Employee {
	var e Employee
	e.Name = name
	e.Salary = 66666
	return &e
}

func printEmployee(e *Employee) {
	fmt.Printf("**********%v**********\n", e.Name)
	fmt.Println("Name: ", e.Name)
	fmt.Println("Rate: ", e.Salary)
	fmt.Println("Address: ", e.Address)
	fmt.Println("Street: ", e.Street)
	fmt.Println("City:   ", e.City)
	fmt.Println("State:  ", e.State)
	fmt.Println("PostalCode: ", e.PostalCode)
	
}

func madeAdress(s string) *Address {
	a := Address{Street : s, City : "Hoston", State: "NE", PostalCode: "714011"}
	return &a
}

func printAddress(a *Address) {
	fmt.Printf("**********%v**********\n", a.Street)
	fmt.Println("Street: ", a.Street)
	fmt.Println("City:   ", a.City)
	fmt.Println("State:  ", a.State)
	fmt.Println("PostalCode: ", a.PostalCode)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m * 3.785)
}