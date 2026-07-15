package belajar_golang_routine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	// Mencoba memperbanyak goroutine
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}
	// Mengambil jumlah CPU yang berada di PC
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	// Mengetahu Total Thread yang sedang berjalan.
	// parameter -1 untuk mengambil, jika diatas 0 digunakan untuk mengubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Mengetahui jumlah Goroutine yang sedang berjalan
	// Secara bawaan akan ada minimal 2 goroutine: 1 untuk menjalankan programnya, dan 1 untuk menghapus (garbage collection) program di memori ketika program sudah selesai dijalankan
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	// Mencoba memperbanyak goroutine
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}
	// Mengambil jumlah CPU yang berada di PC
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	// Menambahkan Jumlah Thread
	// Jarang sekali digunakan, karena golang sudah teroptimalisasi
	runtime.GOMAXPROCS(20)
	// Mengetahu Total Thread yang sedang berjalan.
	// parameter -1 untuk mengambil, jika diatas 0 digunakan untuk mengubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Mengetahui jumlah Goroutine yang sedang berjalan
	// Secara bawaan akan ada minimal 2 goroutine: 1 untuk menjalankan programnya, dan 1 untuk menghapus (garbage collection) program di memori ketika program sudah selesai dijalankan
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}
