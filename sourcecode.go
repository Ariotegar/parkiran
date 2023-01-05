package main

import (
	"fmt"
	"os"
	"os/exec"
)

const T int = 20
const N int = 20
const M int = 20

type kendaraan struct {
	nk          string
	jenis       string
	jam_masuk   int
	menit_masuk int
}

type arrMotor struct {
	data   [T]kendaraan
	jumlah int
}

type arrMobil struct {
	data   [N][M]kendaraan
	jumlah int
}

type struk struct {
	data   kendaraan
	keluar string
}

var area1 arrMotor
var area2 arrMobil

func menu(pilih *int) {
	fmt.Println("SELAMAT DATANG DI PARKIRAN PEKALONGAN, KABUPATEN SPINBOT [v0.9]\nSILAHKAN INPUT NOMOR DIBAWAH INI\n1.  Masukkan data kendaraan\n2.  Lihat parkir motor\n3.  Lihat parkir mobil\n4.  Mengeluarkan kendaraan dan mencatat didalam text file\n5.  Sorting kendaraan sesuai jenis yang diberikan\n6.  Mencari kendaraan menggunakan nomor plat\n7.  Menghitung kendaraan\n8.  Merapihkan parkir motor\n9.  Total penghasilan hari ini\n10. Cek okupansi parkir\n0. exit (selesai)")

	fmt.Scan(&*pilih)
}

func printArea1() {
	var kosong bool = false
	for i := 0; i < T; i++ {
		if area1.data[i].nk == "" {
			fmt.Print("{}", " ")
			kosong = true
		} else {
			fmt.Print(area1.data[i].nk, " ")
			fmt.Print(area1.data[i].jenis, " ")
			fmt.Print(area1.data[i].jam_masuk, ":")
			fmt.Print(area1.data[i].menit_masuk, " ")
		}
	}
	if kosong {
		fmt.Println("\nAda Parkiran Kosong")
	}
}

func printArea2() {
	var kosong bool = false
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if area2.data[i][j].nk == "" {
				fmt.Print("0", " ")
				kosong = true
			} else {
				fmt.Print(area2.data[i][j].nk, " ")
				fmt.Print(area2.data[i][j].jenis, " ")
				fmt.Print(area2.data[i][j].jam_masuk, ":")
				fmt.Print(area2.data[i][j].menit_masuk, " ")
			}
		}
		fmt.Println()
	}
	if kosong {
		fmt.Println("\nAda Parkiran Kosong")
	} else {
		fmt.Println("\nParkiran Penuh")
	}
}

func kendaraanMasuk() {
	var data kendaraan
	var jamvalid bool = false
	var menitvalid bool = false

	fmt.Print("Masukkan plat kendaraan : ")
	fmt.Scan(&data.nk)
	fmt.Print("Masukkan jenis kendaraan : ")
	fmt.Scan(&data.jenis)
	fmt.Print("Masukkan jam masuk : ")
	for !jamvalid {
		fmt.Scan(&data.jam_masuk)
		if data.jam_masuk >= 9 && data.jam_masuk <= 22 {
			jamvalid = true
		} else {
			fmt.Println("Masukan jam yang valid")
		}
	}
	fmt.Print("Masukkan menit masuk : ")
	for !menitvalid {
		fmt.Scan(&data.menit_masuk)
		if data.menit_masuk >= 0 && data.menit_masuk < 60 {
			menitvalid = true
		} else {
			fmt.Println("Masukan menit yang valid")
		}
	}
	var inputed bool = false
	if data.jenis == "motor" {
		for i := 0; i < T && !inputed; i++ {
			if area1.data[i].nk == "" {
				area1.data[i] = data

				area1.jumlah++

				inputed = true

				fmt.Println("Data Berhasil Diinput")
				printArea1()
				fmt.Println("Lokasi kendaraan anda ada di petak :", i)
			}
		}
	} else if data.jenis == "mobil" {

		for i := 0; i < N && !inputed; i++ {
			for j := 0; j < M && !inputed; j++ {
				if area2.data[i][j].nk == "" {
					area2.data[i][j] = data

					area2.jumlah++

					inputed = true

					fmt.Println("Data Berhasil Diinput")

					printArea2()
					fmt.Println("Lokasi kendaraan anda ada di petak :", i+1, ",", j+1)
				}
			}
		}

	} else {

		for i := 0; i < N && !inputed; i++ {
			for j := 0; j < M-1 && !inputed; j++ {
				if area2.data[i][j].nk == "" && area2.data[i][j+1].nk == "" {
					area2.data[i][j] = data
					area2.data[i][j+1] = data

					area2.jumlah += 2

					inputed = true

					fmt.Println("Data Berhasil Diinput")

					printArea2()
					fmt.Println("Lokasi kendaraan anda ada di petak :", i+1, ",", j+1, "dan", i+1, ",", j+2)
				}
			}
		}
	}
}

func mencari() {
	var jenis, plat string
	var i int = 0
	var j int = 0
	var found bool = false
	fmt.Print("Apa jenis kendaraan anda?\nPilihan: (Motor, Mobil, Lainnya)\nPilihan anda : ")
	fmt.Scan(&jenis)
	fmt.Print("Apa nomor plat anda? : ")
	fmt.Scan(&plat)
	if jenis == "Motor" || jenis == "motor" {
		for i < T && !found {
			if area1.data[i].nk == plat {
				fmt.Println("Kendaraan Ditemukan")
				fmt.Println("Plat Nomor : ", area1.data[i].nk, " ")
				fmt.Println("Jenis Kendaraan : ", area1.data[i].jenis, " ")
				fmt.Print("Waktu Masuk : ", area1.data[i].jam_masuk, ":")
				fmt.Print(area1.data[i].menit_masuk, " ")
				found = true
			} else {
				i++
			}
		}
	} else {
		for i < N && !found {
			for j < M && !found {
				if area2.data[i][j].nk == plat {
					fmt.Println("Kendaraan Ditemukan")
					fmt.Println("Plat Nomor : ", area2.data[i][j].nk, " ")
					fmt.Println("Jenis Kendaraan : ", area2.data[i][j].jenis, " ")
					fmt.Print("Waktu Masuk : ", area2.data[i][j].jam_masuk, ":")
					fmt.Print(area2.data[i][j].menit_masuk, " ")
					found = true
				} else {
					j++
				}
			}
			i++
		}
	}
	if !found {
		fmt.Println()
		fmt.Println("Kendaraan tidak ditemukan, kembali lagi ke menu utama..")
	}
	fmt.Println()
	fmt.Println()
}

func sorting() {
	var kendara string
	var i, j, x int
	var arraymotor [T]kendaraan = area1.data
	var arraymobil [N][M]kendaraan
	var arraylain [N][M]kendaraan
	fmt.Println("Jenis Kendaraan yang mau disorting: ")
	fmt.Println("(Motor, Mobil, Lainnya)")
	fmt.Println("DISCLAIMER: SORTING BERDASARKAN STRING")
	fmt.Print("Pilihan :")
	fmt.Scan(&kendara)
	if kendara == "motor" || kendara == "Motor" {
		i = 1
		for i < area1.jumlah {
			j = i - 1
			for j >= 0 && arraymotor[j].nk > arraymotor[j+1].nk {
				tmp := arraymotor[j]
				arraymotor[j] = arraymotor[j+1]
				arraymotor[j+1] = tmp
				j--
			}
			i++
		}
		fmt.Println()
		i = 0
		for i < N {
			if arraymotor[i].nk != "" {
				fmt.Print(arraymotor[i], " ")
				fmt.Println()
			}
			i++
		}
	} else if kendara == "mobil" || kendara == "Mobil" {
		kendara = "mobil"

		i = 0
		for i < N {
			j = 0
			for j < M {
				if area2.data[i][j].jenis == "mobil" {
					arraymobil[i][j] = area2.data[i][j]
				}
				j++
			}
			i++
		}
		i = 1
		for i < area2.jumlah {
			j = i - 1
			for j >= 0 {
				x = i - 1
				for x >= 0 {
					if arraymobil[j][x].nk > arraymobil[j][x+1].nk {
						tmp := arraymobil[j][x]
						arraymobil[j][x] = arraymobil[j][x+1]
						arraymobil[j][x+1] = tmp

						x--
					} else {
						x--
					}
				}
				if arraymobil[j][M-1].nk > arraymobil[j+1][x+1].nk {
					tmp := arraymobil[j][M-1]
					arraymobil[j][M-1] = arraymobil[j+1][x+1]
					arraymobil[j+1][x+1] = tmp
					j--
				} else {
					j--
				}
			}
			i++
		}
		fmt.Println()
		i = 0
		for i < N {
			j = 0
			for j < M {
				if arraymobil[i][j].nk != "" {
					fmt.Print(arraymobil[i][j], " ")
					fmt.Println()
				}
				j++
			}
			i++
		}
	} else if kendara == "Lainnya" || kendara == "lainnya" {
		kendara = "mobil"
		i = 0
		for i < N {
			j = 0
			for j < M {
				if area2.data[i][j].jenis != "mobil" {
					arraylain[i][j] = area2.data[i][j]
				}
				j++
			}
			i++
		}
		i = 1
		for i < area2.jumlah {
			j = i - 1
			for j >= 0 {
				x = i - 1
				for x >= 0 {
					if arraylain[j][x].nk > arraylain[j][x+1].nk {
						tmp := arraylain[j][x]
						arraylain[j][x] = arraylain[j][x+1]
						arraylain[j][x+1] = tmp

						x--
					} else {
						x--
					}
				}
				if arraylain[j][M-1].nk > arraylain[j+1][x+1].nk {
					tmp := arraylain[j][M-1]
					arraylain[j][M-1] = arraylain[j+1][x+1]
					arraylain[j+1][x+1] = tmp
					j--
				} else {
					j--
				}
			}
			i++
		}
		fmt.Println()
		i = 0
		for i < N {
			j = 0
			for j < M {
				if arraylain[i][j].nk != "" {
					fmt.Print(arraylain[i][j], " ")
					fmt.Println()
				}
				j++
			}
			i++
		}
	} else {
		fmt.Println("Input tidak valid")
	}
}

func rapih() {
	var temp kendaraan
	for pass := 0; pass < T-1; pass++ {
		for i := 0; i < T-1; i++ {
			if area1.data[i].nk == "" {
				temp = area1.data[i]
				area1.data[i] = area1.data[i+1]
				area1.data[i+1] = temp
			}
		}
	}
	printArea1()
}

func hitungkendaraan() {
	var i int = 0
	var j int = 0
	var jummotor, jummobil, jumlain int
	for i < T {
		if area1.data[i].jenis == "motor" || area1.data[i].jenis == "Motor" {
			jummotor++
		}
		i++
	}
	i = 0
	for i < N {
		for j < M {
			if area2.data[i][j].jenis == "mobil" || area2.data[i][j].jenis == "Mobil" {
				jummobil++
			} else if area2.data[i][j].jenis != "" {
				jumlain++
			}
			j++
		}
		i++
	}
	fmt.Println("Jumlah Motor yang parkir :", jummotor)
	fmt.Println("Jumlah Mobil kecil yang parkir :", jummobil)
	fmt.Println("Jumlah Kendaraan Lainnya yang parkir :", jumlain/2)
	fmt.Println()
}

func keluar(total *int) {
	var i, j, menitmasuk, menitkeluar, jammasuk, jamkeluar, hargatotal int
	var jenis, plat string
	var valid bool
	valid = false
	fmt.Print("Jenis Kendaraan anda : ")
	fmt.Scan(&jenis)
	fmt.Print("Plat Nomor: ")
	if jenis == "motor" || jenis == "Motor" {
		fmt.Scan(&plat)
		for i < T {
			if area1.data[i].nk == plat {
				fmt.Println("Kendaraan ditemukan, silahkan input jam dan menit keluar")
				jammasuk = area1.data[i].jam_masuk
				menitmasuk = area1.data[i].menit_masuk
				for !valid {
					fmt.Print("Jam: ")
					fmt.Scan(&jamkeluar)
					fmt.Print("Menit: ")
					fmt.Scan(&menitkeluar)
					if jamkeluar > jammasuk {
						valid = true
					} else if jamkeluar == jammasuk {
						if menitkeluar > menitmasuk {
							valid = true
						}
					}
					if !valid {
						fmt.Println("Masukkan tidak valid, mohon input lagi.")
						fmt.Println()
					}
				}
				hargatotal = harga(jammasuk, jamkeluar, menitmasuk, menitkeluar)
				bayar(hargatotal, total)
				write(area1.data[i].nk, jammasuk, menitmasuk, jamkeluar, menitkeluar)
				area1.data[i].nk = ""
				area1.data[i].jenis = ""
				area1.data[i].jam_masuk = 00
				area1.data[i].menit_masuk = 00
				area1.jumlah--
			}
			i++
		}
	} else if jenis == "mobil" || jenis == "Mobil" {
		fmt.Scan(&plat)
		for i < N {
			for j < M {
				if area2.data[i][j].nk == plat {
					fmt.Println("Kendaraan ditemukan, silahkan input jam dan menit keluar")
					jammasuk = area2.data[i][j].jam_masuk
					menitmasuk = area2.data[i][j].menit_masuk
					for !valid {
						fmt.Print("Jam: ")
						fmt.Scan(&jamkeluar)
						fmt.Print("Menit: ")
						fmt.Scan(&menitkeluar)
						if jamkeluar > jammasuk {
							valid = true
						} else if jamkeluar == jammasuk {
							if menitkeluar > menitmasuk {
								valid = true
							}
						}
						if !valid {
							fmt.Println("Masukkan tidak valid, mohon input lagi.")
							fmt.Println()
						}
					}
					hargatotal = harga(jammasuk, jamkeluar, menitmasuk, menitkeluar)
					bayar(hargatotal, total)
					write(area2.data[i][j].nk, jammasuk, menitmasuk, jamkeluar, menitkeluar)
					area2.data[i][j].nk = ""
					area2.data[i][j].jenis = ""
					area2.data[i][j].jam_masuk = 00
					area2.data[i][j].menit_masuk = 00
					area2.jumlah--
				}
				j++
			}
			i++
		}
	} else {
		fmt.Scan(&plat)
		for i < N {
			for j < M {
				if area2.data[i][j].nk == plat && area2.data[i][j+1].nk == plat {
					fmt.Println("Kendaraan ditemukan, silahkan input jam dan menit keluar")
					fmt.Println("Jam: ")
					fmt.Scan(&jamkeluar)
					fmt.Println("Menit: ")
					fmt.Scan(&menitkeluar)
					jammasuk = area2.data[i][j].jam_masuk
					menitmasuk = area2.data[i][j].menit_masuk
					hargatotal = harga(jammasuk, jamkeluar, menitmasuk, menitkeluar)
					bayar(hargatotal, total)
					write(area2.data[i][j].nk, jammasuk, menitmasuk, jamkeluar, menitkeluar)
					area2.data[i][j].nk = ""
					area2.data[i][j].jenis = ""
					area2.data[i][j].jam_masuk = 00
					area2.data[i][j].menit_masuk = 00
					area2.data[i][j+1].nk = ""
					area2.data[i][j+1].jenis = ""
					area2.data[i][j+1].jam_masuk = 00
					area2.data[i][j+1].menit_masuk = 00
					area2.jumlah = area2.jumlah - 2
				}
				j++
			}
			i++
		}
	}
}

func bayar(hargatotal int, total *int) {
	var bayartotal int
	fmt.Println("Total yang harus dibayar: ", hargatotal)
	fmt.Println("Total yang dibayar: ")
	fmt.Scan(&bayartotal)
	for bayartotal < hargatotal {
		fmt.Println("Maaf, total bayar yang anda berikan tidak mencukupi total yang harus dibayar, mohon dicoba lagi.")
		fmt.Println("Total yang harus dibayar: ", hargatotal)
		fmt.Println("Total yang dibayar: ")
		fmt.Scan(&bayartotal)
	}
	fmt.Println("Terimakasih sudah parkir di tempat kami, kembalian anda: ", bayartotal-hargatotal)
	fmt.Println()
	*total = *total + hargatotal
}

func harga(jammasuk, jamkeluar, menitmasuk, menitkeluar int) int {
	var harga, menittotal, jamtotal int
	menittotal = menitkeluar - menitmasuk
	jamtotal = jamkeluar - jammasuk
	if menittotal >= 0 && menittotal <= 10 {
		jamtotal++
	}
	harga = jamtotal * 10000
	return harga
}

func okupansi() {
	var pilih string
	fmt.Print("Masukan Area [AREA 1 = Motor, AREA 2 = Mobil/Lainnya] (ANGKA) : ")
	fmt.Scan(&pilih)
	if pilih == "1" {
		fmt.Println((float64(area1.jumlah)/float64(T))*float64(100), "%")
	} else if pilih == "2" {
		fmt.Println((float64(area2.jumlah)/float64(N*M))*float64(100), "%")
	} else {
		fmt.Println("Pilihan tidak valid, kembali ke main menu..")
	}

}

func write(nomor string, jamin, menitin, hourout, minout int) {
	f, err := os.OpenFile("tiket.txt", os.O_APPEND|os.O_CREATE, 0600) //if FOUND = append file. if !FOUND create file.

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for i := 0; i < 1; i++ {
		_, err = f.WriteString(fmt.Sprintf("nomor kendaraan = %s\njam masuk kendaraan = %02d : %02d\njam keluar kendaraan = %02d : %02d \n\n", nomor, jamin, menitin, hourout, minout)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}

func main() {
	var running bool = true
	var total int = 0

	area1.jumlah = 0
	area2.jumlah = 0

	f, err := os.OpenFile("tiket.txt", os.O_APPEND|os.O_CREATE, 0600)

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()

	for running {

		var pilih int
		menu(&pilih)

		switch pilih {
		case 1:
			// fmt.Print("kendaraan masuk")
			kendaraanMasuk()
		case 2:
			printArea1()
		case 3:
			printArea2()
		case 4:
			keluar(&total)
		case 5:
			sorting()
		case 6:
			mencari()
		case 7:
			hitungkendaraan()
		case 8:
			rapih()
		case 9:
			fmt.Println("Rp.", total)
		case 10:
			okupansi()
		case 0:
			fmt.Println("exiting program..")
			fmt.Println("Terimakasih telah menggunakan program kami")
			fmt.Println("Dibuat Oleh: ")
			fmt.Println("MURSYID NAJIB MUHANA | 1301210411")
			fmt.Println("MUHAMMAD SULTAN NURROCHMAN  | 1301213237")
			fmt.Println("ARIOTEGAR RAMADHANTO PUTRA | 1301213441")
			fmt.Println("KELAS : IF-45-01")
			running = false
		default:
			fmt.Print("Input Nomor yang valid")
			fmt.Println()
			fmt.Println()
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
