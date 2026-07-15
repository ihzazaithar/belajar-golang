package main

import (
	"fmt"
	"time"
)

func main() {

	var waktu time.Time = time.Now()

	fmt.Println(waktu)
	fmt.Println(waktu.Local())

	var waktuUTC time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)

	fmt.Println(waktuUTC)
	fmt.Println(waktuUTC.Local())

	formatter := "2006-01-02 15:04:05" // year-month-day hour:minute:second
	value := "2025-05-29 16:05:10"
	// value := "Salah" // Akan mengakibatkan error

	x, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}

	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2 // Hasil akhirnya dalam milisecond

	fmt.Printf("Waktu: %d", duration3) // Biasanya digunakan dalam implementasi database.

}
