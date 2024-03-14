package main

import "fmt"

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var person Person

	person.nama = "Aurellia"
	person.alamat = "Jakarta"
	person.pekerjaan = "Programmer"
	person.alasan = "Belajar Golang karena ingin mengetahui bahasa Golang"

	fmt.Println(person.nama)
	fmt.Println(person.alamat)
	fmt.Println(person.pekerjaan)
	fmt.Println(person.alasan)
}
