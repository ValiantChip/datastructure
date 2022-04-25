package arraylist_test

import (
	"testing"

	"github.com/ValiantChip/datastructure/arraylist"
)

func TestArrayList(t *testing.T) {

	list := arraylist.NewArrayList[int]()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.Remove(2)
	list.Set(5, 3)

	result := list.Get(3)

	if result != 5 {
		t.Errorf("wanted %d, got %d", 5, result)
	}
}
