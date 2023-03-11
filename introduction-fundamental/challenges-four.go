package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type TemanList interface {
	GetTemanByAbsen(absen int) (*Teman, error)
}

type TemanListImpl struct {
	Temans []*Teman
}

func (tl *TemanListImpl) GetTemanByAbsen(absen int) (*Teman, error) {
	if absen < 1 || absen > len(tl.Temans) {
		return nil, fmt.Errorf("invalid absen number")
	}
	return tl.Temans[absen-1], nil
}

var temanList = &TemanListImpl{
	Temans: []*Teman{
		&Teman{"Ahmad Julianto", "Jakarta", "Mahasiswa", "Switch Career dari designer"},
		&Teman{"Abdur Rahman", "Bandung", "Mahasiswa", "Tertarik dengan Golang"},
		&Teman{"Kelvin Jonathan", "Surabaya", "Mahasiswa", "Bidang Pekerjaan yang mengharuskan pakai golang"},
		&Teman{"Alfia Manurung", "Medan", "Mahasiswa", "Baru lulus, dan baru belajar golang"},
		&Teman{"Deddy Suherman", "Semarang", "Mahasiswa", "Switch Career dari multimedia"},
	},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Command: go run main.go <absen number>")
		return
	}

	absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid absen number")
		return
	}

	teman, err := temanList.GetTemanByAbsen(absen)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}
