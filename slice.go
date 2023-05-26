package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "ganti"
	fmt.Println(slice1)

	months[0] = "mei lagi"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "abil")
	fmt.Println(slice3)

	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)

	fmt.Println(months)

	newslice := make([]string, 2, 5)

	newslice[0] = "abil"
	newslice[1] = "bilhuda"
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copySlice := make([]string, len(newslice), cap(newslice))

	copy(copySlice, newslice)
	fmt.Println(copySlice)

	var iniArray = [5]int{1, 2, 3, 4, 5}
	var iniSlice = []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
