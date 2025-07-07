package helper

import (
	"fmt"
	"testing"
)

// go test -v -run [nama tes],untuk menjalankan test yang di inginkan.

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

// subtest
// go test -v -run 'TestSubTestAdd/AddPertama',menjalankan sesuai dengan nama subtest yang di inginkan
func TestSubTest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Dengan nama Fulan",
			input:    "Fulan",
			expected: "Hi Fulan",
		},
		{
			name:     "Dengan nama kosong",
			input:    "",
			expected: "Hi ",
		},
		{
			name:     "Dengan nama Dina",
			input:    "Dina",
			expected: "Hi Dina",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.input)
			if result != tt.expected {
				t.Errorf("HelloWorld(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// bechmark
// go test -bench=. ,untuk menjalankannya dengan unit test
// go test -bench=. -run=^$ ,untuk menjalankan tanpa unit test (-run=^$ → skip semua unit test (karena regex ^$ = string kosong, jadi gak ada nama test yang cocok))
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("John")
	}
}

// sub benchmark,go test -v -bench=BencmarkSub /nama sub
func BenchmarkSub(b *testing.B) {
	//
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})

	b.Run("Cindi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Cindi")
		}
	})
}

// table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Doni",
			request: "Doni",
		},
		{
			name:    "Eni",
			request: "Eni",
		},
		{
			name:    "Fania Seriosa",
			request: "Fania Seriosa",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
