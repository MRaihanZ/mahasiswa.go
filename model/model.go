package model

import (
	"fmt"
	"m11/entities"
)

func Menu() int{
	var total int
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println("|          Data Mahasiswa          |")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Print("Masukkan Jumlah Mahasiswa : ")
	fmt.Scanln(&total)
	fmt.Println()
	println("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	fmt.Println()
	return total
}

func Data(jumlah int) []entities.Mahasiswa{
	var students []entities.Mahasiswa
	for i := 0; i < jumlah; i++ {
		var student entities.Mahasiswa
		fmt.Println("====================================")
		fmt.Printf("|   Masukkan Data Mahasiswa ke-%d   |\n", i+1)
		fmt.Println("====================================")
		fmt.Print("NPM  : ")
		fmt.Scanln(&student.NPM)
		fmt.Print("Nama : ")
		fmt.Scanln(&student.Nama)
		fmt.Print("UTS  : ")
		fmt.Scanln(&student.UTS)
		fmt.Print("UAS  : ")
		fmt.Scanln(&student.UAS)
		fmt.Println()
		student.Grade = hitungGrade(student.UTS, student.UAS)
		students = append(students, student)
	}
	println("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	fmt.Println()
	return students
}

func hitungGrade(uts, uas float64) string {
	hasil := (uts * 0.6) + (uas * 0.4)
	if hasil >= 80 {return "A"}
	if hasil >= 70 {return "B"}
	if hasil >= 60 {return "C"}
	if hasil >= 50 {return "D"}
	return "E"
}

func DisplayData(students []entities.Mahasiswa) {
	for i, student := range students {
		fmt.Println("============================")
		fmt.Printf("|   Data Mahasiswa ke-%d    |\n", i+1)
		fmt.Println("============================")
		fmt.Printf("NPM   : %s\n", student.NPM)
		fmt.Printf("Nama  : %s\n", student.Nama)
		fmt.Printf("UTS   : %.2f\n", student.UTS)
		fmt.Printf("UAS   : %.2f\n", student.UAS)
		fmt.Printf("Grade : %s\n", student.Grade)
		fmt.Println()
	}
}

func Credit() {
	fmt.Println("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	fmt.Println("Made By: Muhammad Raihan Zhafran (51422127) - 2IA27")
	fmt.Println("Repository: https://github.com/MRaihanZ/mahasiswa.go")
}