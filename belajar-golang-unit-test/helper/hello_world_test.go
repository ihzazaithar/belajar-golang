package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tabel Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ihza",
			request: "Ihza",
		},
		{
			name:    "Izatar",
			request: "Izatar",
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

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Ihza", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ihza")
		}
	})
	b.Run("Izatar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Izatar")
		}
	})
}

// Benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ihza")
	}
}

func BenchmarkHelloWorldIzatar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Izatar")
	}
}

// Konsep Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		nama     string
		request  string
		expected string
	}{
		{
			nama:     "Ihza",
			request:  "Ihza",
			expected: "Hello Ihza",
		},
		{
			nama:     "Izatar",
			request:  "Izatar",
			expected: "Hello Izatar",
		},
		{
			nama:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}
	for _, test := range tests {
		t.Run(test.nama, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Konsep Subtest
func TestSubTest(t *testing.T) {
	t.Run("Ihza", func(t *testing.T) {
		result := HelloWorld("Ihza")
		require.Equal(t, "Hello Ihza", result, "Result must be 'Hello Ihza'")
		// fmt.Println("TestHelloWorld with Require Done")
	})

	t.Run("Izatar", func(t *testing.T) {
		result := HelloWorld("Izatar")
		require.Equal(t, "Hello Izatar", result, "Result must be 'Hello Izatar'")
		// fmt.Println("TestHelloWorld with Require Done")
	})
}

func TestMain(m *testing.M) {
	// Before Unit Test
	fmt.Println("BEFORE UNIT TEST")
	// Kode-Kode Apapun sebelum diTest
	m.Run()

	// After Unit Test
	fmt.Println("AFTER UNIT TEST")
	// Kode-Kode apapun setelah diTest
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}
	result := HelloWorld("Ihza")
	require.Equal(t, "Hello Ihza", result, "Result must be 'Hello Ihza'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ihza")
	require.Equal(t, "Hello Ihza", result, "Result must be 'Hello Ihza'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ihza")
	assert.Equal(t, "Hello Ihza", result, "Result must be 'Hello Ihza'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ihza")

	if result != "Hello Ihza" {
		// unit test failed
		// panic("Result is not Hello Ihza")
		// t.Fail()
		t.Error("Result must be 'Hello Ihza'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldIzatar(t *testing.T) {
	result := HelloWorld("Izatar")

	if result != "Hello Ihza" {
		// unit test failed
		// panic("Result is not Hello Ihza")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Ihza'")
	}
	fmt.Println("TestHelloWorldIzatar Done")
}
