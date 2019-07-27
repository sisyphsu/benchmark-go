package lang

import (
	"math/rand"
	"testing"
)

const int16Count = 1000

var int16Arr = make([]int16, int16Count)
var int16Map = make(map[int16]bool)

func init() {
	for i := 0; i < int16Count; i++ {
		v := int16(rand.Int())
		int16Arr[i] = v
		int16Map[v] = true
	}
	SortInt16Arr(int16Arr)
}

/**
当数组长度为16时：
	BenchmarkInt16Binary-12         200000000                5.90 ns/op            0 B/op          0 allocs/op
	BenchmarkInt16Map-12            100000000                19.9 ns/op            0 B/op          0 allocs/op
当数组长度为64时：
	BenchmarkInt16Binary-12         200000000                8.57 ns/op            0 B/op          0 allocs/op
	BenchmarkInt16Map-12            50000000                 26.2 ns/op            0 B/op          0 allocs/op
当数组长度为256时：
	BenchmarkInt16Binary-12         100000000               11.5 ns/op             0 B/op          0 allocs/op
	BenchmarkInt16Map-12            50000000                27.7 ns/op             0 B/op          0 allocs/op
当数组长度为1024时：
	BenchmarkInt16Binary-12         100000000               15.6 ns/op             0 B/op          0 allocs/op
	BenchmarkInt16Map-12            50000000                29.2 ns/op             0 B/op          0 allocs/op
当数组长度为10240时：
	BenchmarkInt16Binary-12         50000000                27.7 ns/op             0 B/op          0 allocs/op
	BenchmarkInt16Map-12            50000000                30.7 ns/op             0 B/op          0 allocs/op
当数组长度为50000时：
	BenchmarkInt16Binary-12         30000000                40.5 ns/op             0 B/op          0 allocs/op
	BenchmarkInt16Map-12            50000000                35.3 ns/op             0 B/op          0 allocs/op
*/

func BenchmarkInt16Binary(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchInt16(int16Arr, int16(i))
	}
}

func BenchmarkInt16Map(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = int16Map[int16(i)]
	}
}
