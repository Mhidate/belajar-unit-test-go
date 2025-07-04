package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// kode ini akan di jalankan sebelum semua test
	fmt.Println(">> sebelum test dijalankan")

	// Jalankan semua test
	m.Run()

	// kode ini akan dijalankan setelah semua test
	fmt.Println(">> setelah test selesai")

}

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		//Menandai gagal tapi baris kode berikutnya akan di eksekusi(kode dalam fungsi)
		t.Fail()
	}
	fmt.Println("TestHelloWorldJohn Done")
}

func TestHelloWorldDow(t *testing.T) {
	result := HelloWorld("Dow")

	if result != "Hello Dow" {
		//Menandai gagal dan langsung stop,artinya baris kode berikutnya tidak akan di eksekusi(kode dalam fungsi)
		t.FailNow()
	}
	fmt.Println("TestHelloWorldDow Done")
}

func TestHelloWorldLorem(t *testing.T) {
	result := HelloWorld("Lorem")

	if result != "Hello Lorem" {
		//Eksekusi test tetap lanjut ke baris berikutnya sampai fungsi test selesai.
		//memanggil t.Fail()
		t.Error("Harusnya Hello Lorem")
	}
	fmt.Println("TestHelloWorldLorem Done")
}

func TestHelloWorldBen(t *testing.T) {
	result := HelloWorld("Ben")

	if result != "Hello Ben" {
		//Kode setelah t.Fatal() tidak akan dijalankan lagi.
		//memanggil t.FailNow()
		t.Fatal("Harusnya Hello Ben")
	}
	fmt.Println("TestHelloWorldBen Done")
}
