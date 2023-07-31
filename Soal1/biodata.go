package main

import (
	"fmt"
	"os"
	"strconv"
)

// buat struct
type TemanKelas struct {
	No        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// buat slice untuk menyimpan data
var daftarTeman = []TemanKelas{
	{1, "Handy", "Jl Gatsu", "Programmer", "untuk scalling up skill"},
	{2, "Vio", "Jl Sesama", "Satpam", "untuk swicth carrier"},
	{3, "Isma", "Jl Juanda", "Karyawan", "untuk naik grade"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run biodata.go [nomor teman]")
		return
	}

	// ambil nomor teman dari argumen
	nomor, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	// tampilkan biodata berdasarkan nomor
	for _, teman := range daftarTeman {
		if teman.No == nomor {
			fmt.Println("Data teman dengan nomor", nomor, "adalah:")
			fmt.Println("Nama:", teman.Nama)
			fmt.Println("Alamat:", teman.Alamat)
			fmt.Println("Pekerjaan:", teman.Pekerjaan)
			fmt.Println("Alasan Memilih Golang:", teman.Alasan)
			return
		}
	}

	fmt.Println("Teman dengan nomor", nomor, "tidak ditemukan.")
}
