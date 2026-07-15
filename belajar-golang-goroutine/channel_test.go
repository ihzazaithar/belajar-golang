package belajar_golang_routine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// /**Membuat Channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // atau ini untuk memastikan channel tertutup.

	// Memberikan data ke Channel
	// channel <- "Ihza"

	// Mengambil data dari Channel
	// data := <-channel
	// data = <- channel // atau kalau sudah dideklarasikan
	// fmt.Println(<- channel) // atau langsung ke parameter

	// Menutup Channel
	// close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ihza Zaidan Afthar"
		fmt.Println("Data berhasil dikirim ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

// Membuat Channel*/

// /**Channel as Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ihza Zaidan Afthar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Channel as Parameter*/

// /**Channel Only In & Out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Ihza Zaidan Afthar"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Channel Only In & Out*/

// /**Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Ihza"   // Masuk ke buffer
	channel <- "Zaidan" // Masuk ke buffer
	channel <- "Afthar" // Masuk ke buffer

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("Selesai")
}

// Buffered Channel*/

// /** Range Channel

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Perulangan ke-", strconv.Itoa(i))
		}
		close(channel) // Pastikan ini dilakukan, jika tidak maka akan infinite loop/deadlock
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// Range Channel */

// /** Select Range Channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for { // Kombinasikan select dengan menggunakan ini jika data yang dikirim/terima tidak diketahui
		select { // Ini hanya akan menampilkan/menghasilkan data yang tercepat dari channel" yang ada
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 { // Kondisi untuk infinite For
			break
		}
	}
}

// Select Range Channel */

// /** Default Select

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for { // Kombinasikan select dengan menggunakan ini jika data yang dikirim/terima tidak diketahui
		select { // Ini hanya akan menampilkan/menghasilkan data yang tercepat dari channel" yang ada
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default: // Ini akan dilakukan sampai ada data yang masuk
			fmt.Println("Menunggu Data")
		}

		if counter == 2 { // Kondisi untuk infinite For
			break
		}
	}
}

// Default Select */
