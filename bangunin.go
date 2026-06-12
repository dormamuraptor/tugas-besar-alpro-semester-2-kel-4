// Data sample:
// 1 Echo Bandung Semen 4.4 089876543212
// 2 Charlie Jakarta Kayu 4.1 081234567890
// 3 Bravo Bogor Besi 3.9 084765326454
// 4 Alpha Tangerang Kerikil 4.8 089764324544
// 5 Delta Serang Baja 3.5 087432452834
// 6 Gold Bandung Batubata 4.0 083467823456
// 7 Hotel Bandung Paku 4.7 087726331957
// 8 Foxtrot Jakarta Seng 3.9 089463776593

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
	fmt.Println("\n ID ||       NAMA       ||      LOKASI      ||    BAHAN    || RATING || KONTAK ")
	for j = 0; j < i; j++ {
		fmt.Printf(" %-7d %-19s %-19s %-13s %-8.1f %s\n", t[j].id, t[j].nama, t[j].lokasi, t[j].bahan, t[j].rating, t[j].kontak)
	}
}
func updateData(t *tabSupply, i int, target string) {
	var choice, m string
	var n int
	fmt.Print("\nKetik 'Nama' atau 'Lokasi': ")
	fmt.Scan(&choice)
	if choice == "Nama" {
		fmt.Print("\nKetik nama yang ingin dicari: ")
		fmt.Scan(&target)
		n = sequentialSearchName(*t, i, target)
	} else if choice == "Lokasi" {
		fmt.Print("\nKetik lokasi yang ingin dicari: ")
		fmt.Scan(&target)
		n = sequentialSearchLocation(*t, i, target)
	}
	if n == -1 {
		fmt.Println("Data tidak ditemukan!")
	} else {
		fmt.Printf("\nData ditemukan di posisi ke-%d!\n", n+1)
		fmt.Printf(" %-7d %-19s %-19s %-13s %-8.1f %s\n", t[n].id, t[n].nama, t[n].lokasi, t[n].bahan, t[n].rating, t[n].kontak)
		fmt.Print("\nKetik 'Nama'/'Lokasi'/'Bahan'/'Kontak'/'Rating': ")
		fmt.Scan(&m)
		switch m {
		case "Nama":
			fmt.Print("\nMasukkan nama baru: ")
			fmt.Scan(&t[n].nama)
		case "Lokasi":
			fmt.Print("\nMasukkan lokasi baru: ")
			fmt.Scan(&t[n].lokasi)
		case "Bahan":
			fmt.Print("\nMasukkan bahan supply baru: ")
			fmt.Scan(&t[n].bahan)
		case "Kontak":
			fmt.Print("\nMasukkan kontak baru: ")
			fmt.Scan(&t[n].kontak)
		case "Rating":
			fmt.Print("\nMasukkan rating baru: ")
			fmt.Scan(&t[n].rating)
		}
		printData(*t, i)
	}
}
func deleteData(t *tabSupply, i *int, target string) {
	var choice string
	var j, n int
	fmt.Print("\nKetik 'Nama' atau 'Lokasi': ")
	fmt.Scan(&choice)
	if choice == "Nama" {
		fmt.Print("\nKetik nama yang ingin dicari: ")
		fmt.Scan(&target)
		n = sequentialSearchName(*t, *i, target)
	} else if choice == "Lokasi" {
		fmt.Print("\nKetik lokasi yang ingin dicari: ")
		fmt.Scan(&target)
		n = sequentialSearchLocation(*t, *i, target)
	}
	if n == -1 {
		fmt.Println("\nData tidak ditemukan!")
	} else {
		fmt.Println("\nData ditemukan di posisi ke-%d! Menghapus keseluruhan data supplier...", n+1)
		for j = n; j < *i-1; j++ {
			t[j] = t[j+1]
		}
	}
	*i = *i - 1
	printData(*t, *i)
}
func factoryReset(t *tabSupply, i *int) {
	var j int
	for j = *i - 1; j >= 0; j-- {
		(*t)[j].id = 0
		(*t)[j].nama = " "
		(*t)[j].lokasi = " "
		(*t)[j].bahan = " "
		(*t)[j].rating = 0.0
		(*t)[j].kontak = " "
		*i--
	}
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
		if t[mid].lokasi == target {
			return mid
		} else if t[mid].lokasi > target {
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
			if asdes == "Descending" {
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
		if asdes == "Descending" {
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
func selectionSortName(t *tabSupply, i int) {
	var pass, idx, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		idx = pass - 1
		for j = pass; j < i; j++ {
			if (*t)[j].nama < (*t)[idx].nama {
				idx = j
			}
		}
		temp = (*t)[pass-1]
		(*t)[pass-1] = (*t)[idx]
		(*t)[idx] = temp
	}
	printData(*t, i)
}
func selectionSortLocation(t *tabSupply, i int) {
	var pass, idx, j int
	var temp supply
	for pass = 1; pass < i; pass++ {
		idx = pass - 1
		for j = pass; j < i; j++ {
			if (*t)[j].lokasi < (*t)[idx].lokasi {
				idx = j
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
	fmt.Print("\nMasukkan lokasi target: ")
	fmt.Scan(&target)
	sum = 0
	for j = 0; j < i; j++ {
		*avg = *avg + t[j].rating
		if t[j].lokasi == target {
			sum++
		}
	}
	fmt.Printf("\nJumlah supplier di %s: %d\n", target, sum)
	fmt.Printf("Rata-rata rating kepuasan mitra: %.1f\n", (*avg)/float64(j))
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
	fmt.Println("Masukkan riwayat ('SELESAI' jika sudah):")

	catatan = ""
	for k < 99 && catatan != "SELESAI" {
		fmt.Printf("Input riwayat ke-%d: ", k+1)
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
	for choice != 6 {
		fmt.Println("\nMenu Utama")
		fmt.Println("0. Lihat Data")
		fmt.Println("1. Kelola Data Supplier (Ubah / Tambah / Hapus)")
		fmt.Println("2. Urutkan Data Supplier")
		fmt.Println("3. Cari Data Supplier")
		fmt.Println("4. Catatan Riwayat Pelayanan")
		fmt.Println("5. Statistik")
		fmt.Println("6. Keluar Aplikasi")
		fmt.Print("\nPilih Menu (0-6): ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			printData(t, i)
		case 1:
			var modif int
			fmt.Println("\n1. Tambah Data Baru")
			fmt.Println("2. Ubah Data Supplier")
			fmt.Println("3. Hapus Data Supplier")
			fmt.Println("4. Hapus Semua Data")
			fmt.Print("\nPilih Menu (1-4): ")
			fmt.Scan(&modif)

			switch modif {
			case 1:
				fmt.Println("\nMasukkan data baru (ID, nama, lokasi, bahan, rating, kontak):")
				inputData(&t, &i)
				printData(t, i)
			case 2:
				updateData(&t, i, target)
			case 3:
				deleteData(&t, &i, target)
			case 4:
				factoryReset(&t, &i)
			default:
				fmt.Println("Pilihan tidak valid!")
			}
		case 2:
			if i == 0 {
				fmt.Println("Belum ada data supplier untuk diurutkan.")
			} else {
				var menuSort int
				var modeAD, modeSI string

				fmt.Println("\n1. Urut Rating Performa (Numerik)")
				fmt.Println("2. Urut Nama Perusahaan (Alfabetik - Khusus Binary Search)")
				fmt.Println("3. Urut Lokasi Kota (Alfabetik - Khusus Binary Search)")
				fmt.Print("\nPilih (1/2/3): ")
				fmt.Scan(&menuSort)

				switch menuSort {
				case 1:
					fmt.Print("Pilih mode urutan (ketik 'Selection' atau 'Insertion'): ")
					fmt.Scan(&modeSI)
					fmt.Print("Pilih susunan urutan (ketik 'Descending' atau 'Ascending'): ")
					fmt.Scan(&modeAD)
					if modeSI == "Selection" {
						selectionSort(&t, i, modeAD)
						fmt.Println("\nData rating berhasil diurutkan!")
					} else if modeSI == "Insertion" {
						insertionSort(&t, i, modeAD)
						fmt.Println("\nData rating berhasil diurutkan!")
					}
				case 2:
					selectionSortName(&t, i)
					fmt.Println("Data nama berhasil diurutkan!")
				case 3:
					selectionSortLocation(&t, i)
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

				fmt.Println("\n1. Cari Berdasarkan Nama (Sequential Search)")
				fmt.Println("2. Cari Berdasarkan Lokasi (Sequential Search)")
				fmt.Println("3. Cari Berdasarkan Nama (Binary Search - Pastikan Sudah Diurutkan!)")
				fmt.Println("4. Cari Berdasarkan Lokasi (Binary Search - Pastikan Sudah Diurutkan!)")
				fmt.Print("\nPilih (1-4): ")
				fmt.Scan(&menuSearch)

				fmt.Print("\nMasukkan data kata kunci yang dicari (contoh: jika pilih 1/3, maka masukkan nama supplier): ")
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
						fmt.Println("\nData supplier tidak ditemukan.")
					} else {
						fmt.Printf("\nData ditemukan pada indeks ke-%d!\n", hasilIdx+1)
						fmt.Println("\n ID ||       NAMA       ||      LOKASI      ||    BAHAN    || RATING || KONTAK ")
						fmt.Printf(" %-7d %-19s %-19s %-13s %-8.1f %s\n", t[hasilIdx].id, t[hasilIdx].nama, t[hasilIdx].lokasi, t[hasilIdx].bahan, t[hasilIdx].rating, t[hasilIdx].kontak)
					}
				} else {
					fmt.Println("Pilihan invalid.")
				}
			}
		case 4:
			if i == 0 {
				fmt.Println("\nBelum ada data supplier.")
			} else {
				var menuRiwayat int

				fmt.Println("\n1. Tambah Catatan Riwayat Baru")
				fmt.Println("2. Lihat Semua Catatan Riwayat")
				fmt.Print("\nPilih (1/2): ")
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
			fmt.Println("Pilihan menu invalid. Coba lagi.")
		}
	}
}
