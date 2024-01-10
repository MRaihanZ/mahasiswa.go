// Muhammad Raihan Zhafran
// 51422127
// 2IA27
// Repo: https://github.com/MRaihanZ/mahasiswa.go
package main

import (
	"fmt"
	"m11/model"
)

func main() {
	var jumlah int
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println("|          Data Mahasiswa          |")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Print("Masukkan Jumlah Mahasiswa : ")
	fmt.Scanln(&jumlah)
	fmt.Println()
	println("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+")
	fmt.Println()
	data := model.Data(jumlah)
	model.DisplayData(data)
}