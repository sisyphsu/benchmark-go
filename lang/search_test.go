package lang

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"sort"
	"testing"
)

const count = 64

var items = make([]*item, count)
var itemHash = make([]int, count)
var itemHashIndex = make([]int, count)
var itemMap = make(map[string]*item)

type intarray []int

func (a intarray) Len() int           { return len(a) }
func (a intarray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intarray) Less(i, j int) bool { return a[i] < a[j] }

type item struct {
	name  string
	score float64
}

func hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}

func init() {
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("%16d", i)
		score := rand.Float64()
		item := item{name, score}
		itemMap[name] = &item
		items[i] = &item
		itemHash[i] = hashcode(name)
	}
	sort.Sort(intarray(itemHash))
	for i := 0; i < count; i++ {
		off := binarySearch(itemHash, hashcode(items[i].name))
		itemHashIndex[off] = i
	}
}

func binarySearch(arr []int, val int) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		middle := (low + high) / 2
		if arr[middle] == val {
			return middle
		}
		if arr[middle] < val {
			low = middle + 1
		}
		if arr[middle] > val {
			high = middle - 1
		}
	}
	return -1
}

func getByMap(name string) *item {
	item := itemMap[name]
	return item
}

func getByIndex(name string) *item {
	off := binarySearch(itemHash, hashcode(name))
	if off < 0 || off >= len(itemHashIndex) {
		return nil
	}
	i := itemHashIndex[off]
	if i < 0 || i >= len(items) {
		return nil
	}
	return items[i]
}

func TestSearch(t *testing.T) {
	name := items[count/2].name
	item1 := getByMap(name)
	item2 := getByIndex(name)
	if item1.name != name || item2.name != name {
		t.Fatal("error")
	}
}

/**
数据量为64
BenchmarkBinarySearch-12          300000              4157 ns/op            1024 B/op         64 allocs/op
BenchmarkMapSearch-12            2000000               745 ns/op               0 B/op          0 allocs/op

数据量为256
BenchmarkBinarySearch-12          100000             23523 ns/op            4096 B/op        256 allocs/op
BenchmarkMapSearch-12             500000              3511 ns/op               0 B/op          0 allocs/op

数据量为1024
BenchmarkBinarySearch-12           10000            116149 ns/op           16384 B/op       1024 allocs/op
BenchmarkMapSearch-12             100000             20021 ns/op               0 B/op          0 allocs/op
*/

func BenchmarkBinarySearch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			_ = getByIndex(items[j].name).score
		}
	}
}

func BenchmarkMapSearch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < count; j++ {
			_ = getByMap(items[j].name).score
		}
	}
}
