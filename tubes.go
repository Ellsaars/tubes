package main

import (
	"fmt"
	"sort"
)

const NMAX = 100

type Negara struct {
	Nama     string
	Emas     int
	Perak    int
	Perunggu int
}

var negara [NMAX]Negara
var jumlahNegara int

func tambahNegara(nama string) {
	if jumlahNegara < NMAX {
		negara[jumlahNegara] = Negara{Nama: nama}
		jumlahNegara++
		fmt.Println("Negara berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah negara telah mencapai batas maksimal.")
	}
}

func hapusNegara(nama string) {
	for i := 0; i < jumlahNegara; i++ {
		if negara[i].Nama == nama {
			copy(negara[i:], negara[i+1:jumlahNegara])
			jumlahNegara--
			fmt.Println("Negara berhasil dihapus.")
			return
		}
	}
	fmt.Println("Negara tidak ditemukan.")
}

func ubahMedali(nama string, emas, perak, perunggu int) {
	for i := 0; i < jumlahNegara; i++ {
		if negara[i].Nama == nama {
			negara[i].Emas += emas
			negara[i].Perak += perak
			negara[i].Perunggu += perunggu
			fmt.Println("Medali berhasil diubah.")
			return
		}
	}
	fmt.Println("Negara tidak ditemukan.")
}

type MedaliSorter struct {
	negara []Negara
}

func (m MedaliSorter) Len() int {
	return len(m.negara)
}

func (m MedaliSorter) Less(i, j int) bool {
	if m.negara[i].Emas != m.negara[j].Emas {
		return m.negara[i].Emas > m.negara[j].Emas
	}
	if m.negara[i].Perak != m.negara[j].Perak {
		return m.negara[i].Perak > m.negara[j].Perak
	}
	return m.negara[i].Perunggu > m.negara[j].Perunggu
}

func (m MedaliSorter) Swap(i, j int) {
	m.negara[i], m.negara[j] = m.negara[j], m.negara[i]
}

func tampilkanPeringkatKeseluruhan() {
	sort.Sort(MedaliSorter{negara[:jumlahNegara]})
	fmt.Println("Peringkat Keseluruhan:")
	for i := 0; i < jumlahNegara; i++ {
		fmt.Printf("%d. %s - Emas: %d, Perak: %d, Perunggu: %d\n", i+1, negara[i].Nama, negara[i].Emas, negara[i].Perak, negara[i].Perunggu)
	}
}

func main() {
	var pilihan int
	var nama string
	var emas, perak, perunggu int

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Negara")
		fmt.Println("2. Hapus Negara")
		fmt.Println("3. Ubah Medali")
		fmt.Println("4. Peringkat Keseluruhan")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			tambahNegara(nama)
		case 2:
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			hapusNegara(nama)
		case 3:
			fmt.Print("Nama Negara: ")
			fmt.Scanln(&nama)
			fmt.Print("Emas: ")
			fmt.Scanln(&emas)
			fmt.Print("Perak: ")
			fmt.Scanln(&perak)
			fmt.Print("Perunggu: ")
			fmt.Scanln(&perunggu)
			ubahMedali(nama, emas, perak, perunggu)
		case 4:
			tampilkanPeringkatKeseluruhan()
		case 5:
			fmt.Println("Terima kasih, sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
		}
	}
}
