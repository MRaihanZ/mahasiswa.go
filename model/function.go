package model

import (
	"fmt"
	"m11/entities"
)

func Data(jumlah int) []entities.Mahasiswa{
	var mahasiswas []entities.Mahasiswa
	for i := 0; i < jumlah; i++ {
		var mahasiswa entities.Mahasiswa
		fmt.Println("====================================")
		fmt.Printf("|   Masukkan Data Mahasiswa ke-%d   |\n", i+1)
		fmt.Println("====================================")
		fmt.Print("NPM  : ")
		fmt.Scanln(&mahasiswa.NPM)
		fmt.Print("Nama : ")
		fmt.Scanln(&mahasiswa.Nama)
		fmt.Print("UTS  : ")
		fmt.Scanln(&mahasiswa.UTS)
		fmt.Print("UAS  : ")
		fmt.Scanln(&mahasiswa.UAS)
		fmt.Println()
		mahasiswa.Grade = hitungGrade(mahasiswa.UTS, mahasiswa.UAS)
		mahasiswas = append(mahasiswas, mahasiswa)
	}
	println("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	fmt.Println()
	return mahasiswas
}

func hitungGrade(uts, uas float64) string {
	hasil := (uts * 0.6) + (uas * 0.4)
	var grade string
	if hasil >= 80 {
		grade = "A"
	} else if hasil >= 70 {
		grade = "B"
	} else if hasil >= 60 {
		grade = "C"
	} else if hasil >= 50 {
		grade = "D"
	} else {
		grade = "E"
	}
	return grade
}

func DisplayData(mahasiswas []entities.Mahasiswa) {
	for i, mhs := range mahasiswas {
		fmt.Println("============================")
		fmt.Printf("|   Data Mahasiswa ke-%d    |\n", i+1)
		fmt.Println("============================")
		fmt.Printf("NPM   : %s\n", mhs.NPM)
		fmt.Printf("Nama  : %s\n", mhs.Nama)
		fmt.Printf("UTS   : %.2f\n", mhs.UTS)
		fmt.Printf("UAS   : %.2f\n", mhs.UAS)
		fmt.Printf("Grade : %s\n", mhs.Grade)
		fmt.Println()
	}
}