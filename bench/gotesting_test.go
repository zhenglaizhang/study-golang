package gotesting

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mary had a little lamb")
	// b.SetParallelism(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mary had a little lamb")
	// b.SetParallelism(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	// b.SetParallelism(8)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}
