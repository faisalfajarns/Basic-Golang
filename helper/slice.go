package helper

import "fmt"

func Slice() {
	//note : 
	//Array [... / 1-seterusnya]int
	//Slice []int


	var data = [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni","Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = data[4:7] 
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// data[5] = "Diubah"

	// fmt.Println(slice1)
	// fmt.Println(data)

	// slice1[1] = "Juni Lagi"

	fmt.Println(data)

	var slice2 = data[2:4]

	fmt.Println("SLICE2::",slice2)


	var slice3 = append(slice2,"Faisal")

	fmt.Println("APPEND SLICE3::",slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println("RUBAH DATA SLICE3::",slice3)
	fmt.Println("HASIL AKHIR SLICE2::",slice2)
	fmt.Println("HASIL AKHIR SEMUA::",data)


	// untuk tipe data slice apabila lengthnya sudah ditentukan dapat dilakukan ngepush data secara langsung sesuai length yang sudah ditetapkan
	// dan selanjutnya penambahan data dapat dilakukan menggunakan append
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Faisal"
	newSlice[1] = "Fajar"
	newSlice = append(newSlice, "Nursaid")

	fmt.Println("NEW SLICE:: ",cap(newSlice), len(newSlice), newSlice)

	//copy
	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)

	fmt.Println(copySlice)


	
}