package array

import (
	"reflect"
	"testing"
)

func assertCorrentMessage(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
func assertCorrentMessageSlice(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumer(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrentMessage(t, got, want)
	})

	t.Run("conllection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
		assertCorrentMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("two slice test", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrentMessageSlice(t, got, want)
	})
	t.Run("single slice", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		want := []int{3}
		assertCorrentMessageSlice(t, got, want)
	})

}
