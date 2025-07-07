package helper

/*
Perintah
go test, Menjalankan semua unit test di package saat ini.
go test, ./...	Menjalankan semua unit test di semua sub-package dalam project.

go test -v,	Menampilkan output test secara detail (verbose).
go test -run=TestAdd ,run <regex>	Hanya menjalankan test yang namanya cocok dengan regex tertentu.
go test -cover,	Menampilkan code coverage test.
go test -coverprofile=coverage,	Menyimpan laporan code coverage ke file.
  Lalu tampilkan di browser:
  go tool cover -html=coverage
go test -bench=.,	Menjalankan benchmark test (fungsi BenchmarkXXX).
go test -count=1,	Memaksa test dijalankan dari awal tanpa cache.
go test -timeout=10s,	Memberi batas waktu maksimum untuk semua test.
parallel n	Menentukan jumlah maksimal goroutine test yang boleh berjalan bersamaan.
 go test -parallel=4
*/

/*
Saat  html terbuka menggunakan go tool cover -html=coverage ,maka akan ada keterangan ini

Not Tracked	Putih / tanpa highlight	Baris kode ini tidak dihitung coverage-nya oleh Go (misal: deklarasi import, type struct, const, dan lain-lain).
Not Covered	Merah / pink	Baris kode ini dieksekusi dalam program, tapi tidak pernah ter-cover oleh test.
Covered	Hijau	Baris kode ini sudah dijalankan waktu unit test dijalankan.
*/

func HelloWorld(name string) string {
	return "Hello " + name
}
