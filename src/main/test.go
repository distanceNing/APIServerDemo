package main

import (
	"fmt"
	"log"
)

// T is for test fmt
type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
func deferFunc() {
	fmt.Println("run defer func ")
}

func testDefer() {
	fmt.Println("before defer")
	defer deferFunc()
	fmt.Println("after defer")
}

func testMap() {
	var t = T{1, 2, "aaa"}

	fmt.Printf("%v\n", t.String())
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	tz := "UTC"
	seconds, ok := timeZone[tz]
	if !ok {
		log.Println("unknown time zone:", tz)
	}
	fmt.Println("tz in the map : ", seconds)

	//删除某个键
	delete(timeZone, "PDT")

	for k, v := range timeZone {
		fmt.Println(k, v)
	}
}
