package main

import (
	"math/rand"
	"testing"
)

const num = 2000

func TestFizzBuzz(t *testing.T) {
	a := FizzBuzz(10)
	if len(a) != 9 {
		t.Error("asdas")
	}
}

/*func BenchmarkInsertSort(b *testing.B) {
	nums := make([]int, num)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < num; i++ {
			nums[i] = rand.Int() % 100
		}
		InsertSort(nums)
	}
}
func BenchmarkMaoPaoSort(b *testing.B) {
	nums := make([]int, num)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < num; i++ {
			nums[i] = rand.Int() % 100
		}
		MaopaoSort(nums)
	}
}
func BenchmarkQuickSort(b *testing.B) {
	nums := make([]int, num)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < num; i++ {
			nums[i] = rand.Int() % 100
		}
		QuickSort(nums, 0, len(nums)-1)
	}
}*/

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := rand.Int() % 2
		y := rand.Int() % 2
		test1(x, y)
	}
}
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := rand.Int() % 2
		y := rand.Int() % 2
		test2(x, y)
	}
}
