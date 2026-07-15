package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// ** Pengenalan Context

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Pengenalan Context *

// ** Membuat Child dengan context.WithValue(parent, key, value)

func TestContextWithValue(t *testing.T) {
	// Inisiasi awal Context
	contextA := context.Background()

	// Menambahkan Value ke Context yang mana membuat Child Context
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "E", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// Mendapatkan Value dari Context
	fmt.Println(contextF.Value("f")) // Dapat, karena value Context F itu sendiri.
	fmt.Println(contextF.Value("c")) // Dapat, Value parent dari Context F
	fmt.Println(contextF.Value("b")) // Tidak dapat, karena Context F dan parentnya tidak memiliki value tersebut.
	fmt.Println(contextA.Value("b")) // Tidak dapat, karena pencariannya adalah ke Diri sendiri atau parentnya. Bukan ke Childnya.
}

// Membuat Child dengan context.WithValue(parent, key, value) *

// ** Signal Cancel

// Goroutine Leak

// func CreateCounter() chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		counter := 1

// 		for {
// 			destination <- counter
// 			counter++
// 		}
// 	}()

// 	return destination
// }

// Goroutine Leak

// Goroutine Leak Mitigation with Cancel Signal

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulasi slow
			}
		}
	}()

	return destination
}

// Goroutine Leak Mitigation with Cancel Signal

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	// Deklarasi Cancel Signal
	contextParent := context.Background()
	ctx, cancel := context.WithCancel(contextParent)

	// destination := CreateCounter()
	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel() // Mengirim sinyal cancel ke Context

	time.Sleep(2 * time.Second) // Buffer untuk asynch

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

// Signal Cancel *

// ** Signal Cancel with Timeout

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	// Deklarasi Cancel Signal
	contextParent := context.Background()
	ctx, cancel := context.WithTimeout(contextParent, 5*time.Second)
	defer cancel() // Ini tetap dipanggil jikalau, proses selesai tidak lebih dari Timeoutnya

	// destination := CreateCounter()
	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) // Buffer untuk asynch

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

// Signal Cancel with Timeout *

// ** Signal Cancel with Deadline

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	// Deklarasi Cancel Signal
	contextParent := context.Background()
	ctx, cancel := context.WithDeadline(contextParent, time.Now().Add(5*time.Second))
	defer cancel() // Ini tetap dipanggil jikalau, proses selesai tidak lebih dari Timeoutnya

	// destination := CreateCounter()
	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) // Buffer untuk asynch

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

}

// Signal Cancel with Deadline *
