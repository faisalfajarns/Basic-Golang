package helper

import "fmt"

func Boolean_Operation() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusAbsensi && lulusNilaiAkhir

	fmt.Println(lulus)
}