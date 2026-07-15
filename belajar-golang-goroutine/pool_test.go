package belajar_golang_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// pool := sync.Pool{}
	// Sintaks dibawah adalah sintaks yang memberikan nilai default ke Pool jika, ketika menggunakan .Get() ternyata nil, maka akan dikembalikan nilai yang ditentukan.
	pool := sync.Pool{
		New: func() interface{} {
			return "New" // Ini adalah nilai default jika Pool kosong.
		},
	}

	pool.Put("Ihza") // Memberikan data ke Pool
	pool.Put("Zaidan")
	pool.Put("Afthar")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // Mengambil data dari Pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")

}
