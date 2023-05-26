package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75
	var lulusUJian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUJian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUJian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

}
