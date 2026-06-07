package main

import (
	"fmt"
)

type supply struct {
	id, nRiwayat                int
	nama, lokasi, bahan, kontak string
	riwayat                     [99]string
	rating                      float64
}
type tabSupply [999]supply

// CRUD (CREATE, READ, UPDATE, DELETE)
func inputData(t *tabSupply, i *int) {
	fmt.Scan(&t[*i].id)
	for t[*i].id != -1 {
		fmt.Scan(&t[*i].nama, &t[*i].lokasi, &t[*i].bahan, &t[*i].rating, &t[*i].kontak)
		*i++
		fmt.Scan(&t[*i].id)
	}
}
func printData(t tabSupply, i int) {
	var j int
	for j = 0; j < i; j++ {
		fmt.Println(t[j].id, t[j].nama, t[j].lokasi, t[j].bahan, t[j].rating, t[j].kontak)
	}
}
func updateData(t *tabSupply, i int, target string) {
	var choice, m string
	var n int
	fmt.Println("pilih nama atau lokasi")
	fmt.Scan(&choice)
	if choice == "nama" {
		n = sequentialSearchName(*t, i, target)
	} else if choice == "lokasi" {
		n = sequentialSearchLocation(*t, i, target)
	}
	fmt.Println("pilih:")
	fmt.Scan(&m)
	switch m {
	case "nama":
		fmt.Println("masukkan nama baru")
		fmt.Scan(&t[n].nama)
	case "lokasi":
		fmt.Println("masukkan lokasi baru")
		fmt.Scan(&t[n].lokasi)
	case "bahan":
		fmt.Println("masukkan bahan supply baru")
		fmt.Scan(&t[n].bahan)
	case "kontak":
		fmt.Println("masukkan kontak baru")
		fmt.Scan(&t[n].kontak)
	case "rating":
		fmt.Println("masukkan rating baru")
		fmt.Scan(&t[n].rating)
	}
	printData(*t, i)
}
func deleteData(t *tabSupply, i *int, target string) {
	var choice string
	var j, n int
	fmt.Println("pilih nama atau lokasi")
	fmt.Scan(&choice)
	if choice == "nama" {
		n = sequentialSearchName(*t, *i, target)
	} else if choice == "lokasi" {
		n = sequentialSearchLocation(*t, *i, target)
	}
	for j = n; j < *i-1; j++ {
		t[j] = t[j+1]
	}
	*i = *i - 1
	printData(*t, *i)
}

// ALGORITMA PENCARIAN
func sequentialSearchName(t tabSupply, i int, target string) int {
	var j int
	for j = 0; j < i; j++ {
		if t[j].nama == target {
			return j
		}
	}
	return -1
}
func sequentialSearchLocation(t tabSupply, i int, target string) int {
	var j int
	for j = 0; j < i; j++ {
		if t[j].lokasi == target {
			return j
		}
	}
	return -1
}
func binarySearchName(t tabSupply, i int, target string) int {
	var left, mid, right int
	left = 0
	right = i - 1
	for left <= right {
		mid = (left + right) / 2
		if t[mid].nama == target {
			return mid
		} else if t[mid].nama > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func binarySearchLocation(t tabSupply, i int, target string) int {
	var left, mid, right int
	left = 0
	right = i - 1
	for left <= right {
		mid = (left + right) / 2
		if t[mid].nama == target {
			return mid
		} else if t[mid].nama > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// ALGORITMA PENGURUTAN (NUMERIK)
func selectionSort(t *tabSupply, i int, asdes string) {
	var pass, idx, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		idx = pass - 1
		for j = pass; j < i; j++ {
			if asdes == "desc" {
				if (*t)[j].rating > (*t)[idx].rating {
					idx = j
				}
			} else {
				if (*t)[j].rating < (*t)[idx].rating {
					idx = j
				}
			}
		}
		temp = (*t)[pass-1]
		(*t)[pass-1] = (*t)[idx]
		(*t)[idx] = temp
	}
	printData(*t, i)
}
func insertionSort(t *tabSupply, i int, asdes string) {
	var pass, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		j = pass
		temp = t[pass]
		if asdes == "desc" {
			for j > 0 && t[j-1].rating < temp.rating {
				t[j] = t[j-1]
				j = j - 1
			}
			t[j] = temp
		} else {
			for j > 0 && t[j-1].rating > temp.rating {
				t[j] = t[j-1]
				j = j - 1
			}
			t[j] = temp
		}
	}
	printData(*t, i)
}

// ALGORITMA PENGURUTAN (ALFABETIK, KHUSUS BINARY SEARCH)
func selectionSortName(t *tabSupply, i int, asdes string) {
	var pass, idx, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		idx = pass - 1
		for j = pass; j < i; j++ {
			if asdes == "desc" {
				if (*t)[j].nama > (*t)[idx].nama {
					idx = j
				}
			} else {
				if (*t)[j].nama < (*t)[idx].nama {
					idx = j
				}
			}
		}
		temp = (*t)[pass-1]
		(*t)[pass-1] = (*t)[idx]
		(*t)[idx] = temp
	}
	printData(*t, i)
}
func selectionSortLocation(t *tabSupply, i int, asdes string) {
	var pass, idx, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		idx = pass - 1
		for j = pass; j < i; j++ {
			if asdes == "desc" {
				if (*t)[j].lokasi > (*t)[idx].lokasi {
					idx = j
				}
			} else {
				if (*t)[j].lokasi < (*t)[idx].lokasi {
					idx = j
				}
			}
		}
		temp = (*t)[pass-1]
		(*t)[pass-1] = (*t)[idx]
		(*t)[idx] = temp
	}
	printData(*t, i)
}

// STATISTIK (SUM & AVG)
func statistic(t tabSupply, i int, target string, avg *float64) {
	var sum, j int
	sum = 0
	for j = 0; j < i; j++ {
		*avg = *avg + t[j].rating
		if t[j].lokasi == target {
			sum++
		}
	}
	fmt.Printf("%d %.1f", sum, *avg/float64(j))
}

// CATAT RIWAYAT SUPPLIER
func history(t *tabSupply, i int, target string) {
	var j, k int
	var catatan string

	j = sequentialSearchName(*t, i, target)
	if j == -1 {
		fmt.Println("Data tidak ditemukan!")
		return
	}
	k = (*t)[j].nRiwayat
	fmt.Println("Masukkan riwayat (SELESAI jika sudah):")

	catatan = ""
	for k < 99 && catatan != "SELESAI" {
		fmt.Printf("Input riwayat ke-%d:", k+1)
		fmt.Scan(&catatan)
		if catatan != "SELESAI" {
			(*t)[j].riwayat[k] = catatan
			(*t)[j].nRiwayat++
			k++
		}
	}
	fmt.Println("Riwayat berhasil disimpan!")
}
func printHistory(t tabSupply, i int, target string) {
	var j, k int

	j = sequentialSearchName(t, i, target)
	if j == -1 {
		fmt.Println("Data supplier tidak ditemukan")
		return
	}
	fmt.Printf("Riwayat Pelayanan: %s\n", t[j].nama)
	for k = 0; k < t[j].nRiwayat; k++ {
		fmt.Println("-", t[j].riwayat[k])
	}
}

func main() {
	var t tabSupply
	var i, choice int
	var target string
	i = 0
	fmt.Println("Masukkan data supplier:")
	inputData(&t, &i)
	printData(t, i)
	for choice != 6 {
		fmt.Println("Menu Utama")
		fmt.Println("1. Kelola Data Supplier (Ubah / Tambah / Hapus)")
		fmt.Println("2. Urutkan Data Supplier")
		fmt.Println("3. Cari Data Supplier")
		fmt.Println("4. Catatan Riwayat Pelayanan")
		fmt.Println("5. Statistik")
		fmt.Println("6. Keluar Aplikasi")
		fmt.Println("Pilih Menu (1-6):")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var modif int
			fmt.Println("1. Tambah Data Baru")
			fmt.Println("2. Ubah Data Supplier")
			fmt.Println("3. Hapus Data Supplier")
			fmt.Print("Pilih (1/2/3):")
			fmt.Scan(&modif)

			switch modif {
			case 1:
				inputData(&t, &i)
			case 2:
				updateData(&t, i, target)
			case 3:
				deleteData(&t, &i, target)
			default:
				fmt.Println("Pilihan tidak valid!")
			}
		case 2:
			if i == 0 {
				fmt.Println("Belum ada data supplier untuk diurutkan.")
			} else {
				var menuSort int
				var modeAD string

				fmt.Println("1. Urut Rating Performa (Numerik)")
				fmt.Println("2. Urut Nama Perusahaan (Alfabetik - Khusus Binary Search)")
				fmt.Println("3. Urut Lokasi Kota (Alfabetik - Khusus Binary Search)")
				fmt.Print("Pilih (1/2/3): ")
				fmt.Scan(&menuSort)

				fmt.Print("Pilih mode (Ketik 'Descending' atau 'Ascending'): ")
				fmt.Scan(&modeAD)

				switch menuSort {
				case 1:
					selectionSort(&t, i, modeAD)
					fmt.Println("Data rating berhasil diurutkan!")
				case 2:
					selectionSortName(&t, i, modeAD)
					fmt.Println("Data nama berhasil diurutkan!")
				case 3:
					selectionSortLocation(&t, i, modeAD)
					fmt.Println("Data lokasi berhasil diurutkan!")
				default:
					fmt.Println("Pilihan tidak valid!")
				}
			}
		case 3:
			if i == 0 {
				fmt.Println("Belum ada data supplier untuk dicari.")
			} else {
				var menuSearch, hasilIdx int

				fmt.Println("1. Cari Berdasarkan Nama (Sequential Search)")
				fmt.Println("2. Cari Berdasarkan Lokasi (Sequential Search)")
				fmt.Println("3. Cari Berdasarkan Nama (Binary Search - Pastikan Sudah Diurutkan!)")
				fmt.Println("4. Cari Berdasarkan Lokasi (Binary Search - Pastikan Sudah Diurutkan!)")
				fmt.Print("Pilih (1-4): ")
				fmt.Scan(&menuSearch)

				fmt.Print("Masukkan data kata kunci yang dicari: ")
				fmt.Scan(&target)

				switch menuSearch {
				case 1:
					hasilIdx = sequentialSearchName(t, i, target)
				case 2:
					hasilIdx = sequentialSearchLocation(t, i, target)
				case 3:
					hasilIdx = binarySearchName(t, i, target)
				case 4:
					hasilIdx = binarySearchLocation(t, i, target)
				default:
					fmt.Println("Pilihan tidak valid!")
				}

				if menuSearch >= 1 && menuSearch <= 4 {
					if hasilIdx == -1 {
						fmt.Println("Data supplier tidak ditemukan.")
					} else {
						fmt.Printf("Data ditemukan pada indeks ke-%d!\n", hasilIdx)
					}
				} else {
					fmt.Println("Pilihan invalid.")
				}
			}
		case 4:
			if i == 0 {
				fmt.Println("Belum ada data supplier.")
			} else {
				var menuRiwayat int

				fmt.Println("1. Tambah Catatan Riwayat Baru")
				fmt.Println("2. Lihat Semua Catatan Riwayat")
				fmt.Print("Pilih (1/2): ")
				fmt.Scan(&menuRiwayat)

				fmt.Print("Masukkan nama perusahaan supplier: ")
				fmt.Scan(&target)

				switch menuRiwayat {
				case 1:
					history(&t, i, target)
				case 2:
					printHistory(t, i, target)
				default:
					fmt.Println("Pilihan tidak valid!")
				}
			}
		case 5:
			if i == 0 {
				fmt.Println("Belum ada data supplier.")
			} else {
				var avg float64
				statistic(t, i, target, &avg)
			}

		case 6:
			fmt.Println("Program selesai. Sampai jumpa!")

		default:
			fmt.Println("Pilihan menu invalid, silahkan coba lagi.")
		}
	}
}
