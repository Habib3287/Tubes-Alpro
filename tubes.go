package main

import (
	"fmt"
)

const NMAX = 100

type dataF struct {
	judul, genre string
	jadwal       string
	durasi       int
	rating       float32
	harga        int
	kursi        int
	promo        bool
}

type film [NMAX]dataF

var GREEN = "\033[1;92m"
var RED = "\033[1;91m"
var RESET = "\033[0;97m"
var LBLUE = "\033[38;2;120;172;255m"

func main() {
	var M film
	var num int
	var n int = 0
	var s string
	var found int

	for {
		fmt.Println(LBLUE + "╔═══════════════════════════════════════════════╗")
		fmt.Println("║ BBBB   IIIII   OOO    SSS  K   K  OOO   PPPP  ║")
		fmt.Println("║ B   B    I    O   O  S     K  K  O   O  P   P ║")
		fmt.Println("║ BBBB     I    O   O  SSSS  KKK   O   O  PPPP  ║")
		fmt.Println("║ B   B    I    O   O     S  K  K  O   O  P     ║")
		fmt.Println("║ BBBB   IIIII   OOO   SSS   K   K  OOO   P     ║")
		fmt.Println("╠═══════════════════════════════════════════════╣")
		fmt.Println("║[1] Pemesanan Tiket                            ║")
		fmt.Println("║[2] Pendataan Film                             ║")
		fmt.Println("║[3] Pencarian Film                             ║")
		fmt.Println("║[4] Edit Film                                  ║")
		fmt.Println("║[5] Exit                                       ║")
		fmt.Println("╚═══════════════════════════════════════════════╝")
		fmt.Println("Pilih: ", RESET)
		fmt.Scan(&num)
		if num == 1 {
			var pilihan int
			fmt.Println(LBLUE + "╔═══════════════════════════════════════════════╗")
			fmt.Println("║Pilih Tampilan                                 ║")
			fmt.Println("║[1] Judul A-Z                                  ║")
			fmt.Println("║[2] Judul Z-A                                  ║")
			fmt.Println("║[3] Jadwal tayang				║")
			fmt.Println("║[4] Rating                              	║")
			fmt.Println("╚═══════════════════════════════════════════════╝")
			fmt.Println("Pilih: ", RESET)
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				sortJudulDes(&M, n)
				cetakData(M, n)
			} else if pilihan == 2 {
				sortJudulAsc(&M, n)
				cetakData(M, n)
			} else if pilihan == 3 {
				sortJadwalTayang(&M, n)
				cetakData(M, n)
			} else if pilihan == 4 {
				sortRating(&M, n)
				cetakData(M, n)
			} else {
				fmt.Println("══╣ Pilihan Tidak Ditemukan ╠══")
			}
			PemesananTiket(&M, n)
		} else if num == 2 {
			pendataanFilm(&M, &n)
		} else if num == 3 {
			fmt.Println(LBLUE + "╔═══════════════════════════════════════════════╗")
			fmt.Println("║Pencarian Film:                                ║")
			fmt.Println("║[1] Berdasarkan judul                          ║")
			fmt.Println("║[2] Berdasarkan genre                          ║")
			fmt.Println("║[3] Berdasarkan tanggal                        ║")
			fmt.Println("╚═══════════════════════════════════════════════╝" + RESET)
			var pilihan int
			fmt.Print(LBLUE+"Pilih: ", RESET)
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				fmt.Println("Masukkan judul film yang dicari:")
				fmt.Scan(&s)
				found = pencarianFilmJudul(M, n, s)
				if found != -1 {
					fmt.Println(LBLUE+"Judul", "Genre", "Jadwal", "Durasi", "Rating", RESET)
					fmt.Println(M[found].judul, M[found].genre, M[found].jadwal, M[found].durasi, M[found].rating)
					fmt.Println(GREEN+"══╣ Film Ditemukan ╠══", RESET)
				} else {
					fmt.Println(RED+"══╣ Film Tidak Ditemukan ╠══", RESET)
				}
			} else if pilihan == 2 {
				fmt.Println("Masukkan genre film yang dicari:")
				fmt.Scan(&s)
				found = pencarianFilmGenre(M, n, s)
			} else if pilihan == 3 {
				fmt.Println("Masukkan tanggal film yang dicari:")
				fmt.Scan(&s)
				found = pencarianFilmTanggal(M, n, s)
			} else {
				fmt.Println(RED+"══╣ Pilihan Tidak Ditemukan ╠══", RESET)
			}
		} else if num == 4 {
			fmt.Println(LBLUE + "╔═══════════════════════════════════════════════╗")
			fmt.Println("║Edit Film                                      ║")
			fmt.Println("║[1] Edit spesifikasi film                      ║")
			fmt.Println("║[2] Penghapusan Film                           ║")
			fmt.Println("╚═══════════════════════════════════════════════╝")
			fmt.Println("Pilih: ", RESET)
			var pil, found int
			fmt.Scan(&pil)
			if pil == 1 {
				fmt.Println("Masukkan judul film yang akan diedit:")
				fmt.Scan(&s)
				found = pencarianFilmJudul(M, n, s)
				if found != -1 {
					var Jadwal string
					var Rating float32
					var Promo bool
					fmt.Print(LBLUE + "Masukkan detail baru")
					fmt.Print("Jadwal: ", RESET)
					fmt.Scan(&Jadwal)
					fmt.Print(LBLUE+"Rating: ", RESET)
					fmt.Scan(&Rating)
					fmt.Print(LBLUE+"Promo: ", RESET)
					fmt.Scan(&Promo)
					editFilm(&M, n, s, Jadwal, Rating, Promo)
					fmt.Println(GREEN+"══╣ Film Berhasil Diubah ╠══", RESET)
				} else {
					fmt.Println(RED+"══╣ Film Tidak Ditemukan ╠══", RESET)
				}
			} else if pil == 2 {
				fmt.Println("Masukkan judul film yang akan dihapus:")
				hapusFilm(&M, &n)
			} else {
				fmt.Println(RED+"══╣ Pilihan Tidak Ditemukan ╠══", RESET)
			}
		} else if num == 5 {
			fmt.Println(LBLUE+"══╣ Terima Kasih ╠══", RESET)
			return
		} else {
			fmt.Println(RED+"══╣ Pilihan Tidak Ditemukan ╠══", RESET)
		}
	}
}

func PemesananTiket(M *film, n int) {
	var judul, x string
	var jumlah, found, total int
	fmt.Println("Judul film yang ingin dibeli: ")
	fmt.Scan(&judul)
	found = pencarianFilmJudul(*M, n, judul)
	if found != -1 {
		fmt.Println("Film tersedia")
		fmt.Println("Berapa tiket yang akan dibeli?")
		fmt.Scan(&jumlah)
		if jumlah > M[found].kursi {
			fmt.Println("Jumlah kursi tidak tersedia")
		} else {
			if M[found].promo == true {
				M[found].harga = (M[found].harga - ((M[found].harga * 25) / 100))
			}
			total = jumlah * M[found].harga
			fmt.Println(LBLUE+"Total harga: ", RESET, total)
			fmt.Println("Konfirmasi pembelian(Y/N)")
			fmt.Scan(&x)
			if x == "Y" || x == "y" {
				fmt.Println(GREEN+"══╣ Pembelian berhasil! ╠══", RESET)
				fmt.Println(LBLUE+"Jumlah tiket: ", RESET, jumlah)
				fmt.Println(LBLUE+"Total harga: ", RESET, total)
			} else {
				fmt.Println(RED+"══╣ Pilih Kembali Film Yang Ingin Dibeli ╠══", RESET)
			}
			M[found].kursi -= jumlah
		}
	} else {
		fmt.Println(RED+"══╣ Tiket yang tersedia tidak mencukupi ╠══", RESET)
	}
}

func cetakData(M film, n int) {
	var i int
	fmt.Println(LBLUE + "════════════════╣ DAFTAR FILM ╠════════════════")
	fmt.Println("No", "Judul", "Genre", "Jadwal", "Durasi", "Rating", "Harga", "Jumlah Kursi", "Promo", RESET)
	for i = 0; i < n; i++ {
		fmt.Println(i+1, ".", M[i].judul, M[i].genre, M[i].jadwal, M[i].durasi, M[i].rating, M[i].harga, M[i].kursi, M[i].promo)
	}
}

func sortJudulAsc(M *film, n int) {
	var i, j int
	var temp dataF

	for i = 1; i < n; i++ {
		temp = M[i]
		j = i
		for j > 0 && temp.judul > M[j-1].judul {
			M[j] = M[j-1]
			j--
		}
		M[j] = temp
	}
}

func sortJudulDes(M *film, n int) {
	var i, j int
	var temp dataF

	for i = 1; i < n; i++ {
		temp = M[i]
		j = i
		for j > 0 && temp.judul < M[j-1].judul {
			M[j] = M[j-1]
			j--
		}
		M[j] = temp
	}
}

func sortJadwalTayang(M *film, n int) {
	var i, j int
	var temp dataF

	for i = 1; i < n; i++ {
		temp = M[i]
		j = i
		for j > 0 && temp.jadwal < M[j-1].jadwal {
			M[j] = M[j-1]
			j--
		}
		M[j] = temp
	}
}

func sortRating(M *film, n int) {
	var i, j int
	var temp dataF

	for i = 1; i < n; i++ {
		temp = M[i]
		j = i
		for j > 0 && temp.rating > M[j-1].rating {
			M[j] = M[j-1]
			j--
		}
		M[j] = temp
	}
}

func pendataanFilm(M *film, n *int) {
	var judul, genre, jadwal string
	var durasi int
	var rating float32
	var harga, kursi int
	var promo bool
	for *n < NMAX {
		fmt.Println(LBLUE+"Judul: (Ketik 'sudah' untuk berhenti)", RESET)
		fmt.Scan(&judul)
		if judul == "sudah" || judul == "Sudah" {
			return
		}
		fmt.Println(LBLUE+"Genre: ", RESET)
		fmt.Scan(&genre)
		fmt.Println(LBLUE+"Jadwal tayang (YYYY-MM-DD): ", RESET)
		fmt.Scan(&jadwal)
		fmt.Println(LBLUE+"Durasi (menit): ", RESET)
		fmt.Scan(&durasi)
		fmt.Println(LBLUE+"Rating: ", RESET)
		fmt.Scan(&rating)
		fmt.Println(LBLUE+"Harga tiket: ", RESET)
		fmt.Scan(&harga)
		fmt.Println(LBLUE+"Jumlah kursi: ", RESET)
		fmt.Scan(&kursi)
		fmt.Println(LBLUE+"Promo: (True/False)", RESET)
		fmt.Scan(&promo)
		M[*n] = dataF{judul, genre, jadwal, durasi, rating, harga, kursi, promo}
		*n += 1
		fmt.Println(GREEN+"══╣ Film Berhasil Ditambahkan ╠══", RESET)
	}
	fmt.Println(RED+"══╣ Daftar Film Sudah Penuh ╠══", RESET)
}

func tampilkanFilm(M film, n int) {
	fmt.Println(M[n].judul, M[n].genre, M[n].jadwal, M[n].durasi, M[n].rating)
	fmt.Println("════════════════════════════════════════════════════════════════════════")
}

func pencarianFilmJudul(M film, n int, s string) int {
	var i int
	for i = 0; i < n; i++ {
		if M[i].judul == s {
			return i
		}
	}

	return -1
}
func pencarianFilmGenre(M film, n int, s string) int {
	var i int
	for i = 0; i < n; i++ {
		if M[i].genre == s {
			tampilkanFilm(M, i)
		}
	}
	return -1
}
func pencarianFilmTanggal(M film, n int, s string) int {
	var i int
	for i = 0; i < n; i++ {
		if M[i].jadwal == s {
			tampilkanFilm(M, i)
		}
	}
	return -1
}

func editFilm(M *film, i int, s string, jadwal string, rating float32, promo bool) {
	(*M)[i].jadwal = jadwal
	(*M)[i].rating = rating
	(*M)[i].promo = promo
}

func hapusFilm(M *film, n *int) {
	var i, found int
	var s string
	fmt.Scan(&s)
	found = pencarianFilmJudul(*M, *n, s)
	if found != -1 {
		for i = found; i <= *n-2; i++ {
			(*M)[i] = (*M)[i+1]
		}
		*n--
		fmt.Println(GREEN+"══╣ Film Berhasil Dihapus ╠══", RESET)
	} else {
		fmt.Println(RED+"══╣ Film Tidak Ditemukan ╠══", RESET)
	}
}
