package doublylinked

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	list := &List{}
	values := [9]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := len(values)
	for _, v := range values {
		list.PushBack(v)
	}

	// get firsr item
	t.Run("First", func(t *testing.T) {
		n := list.First()
		if reflect.ValueOf(n.Value()).Int() != values[0] {
			t.Errorf("first item should be %d", values[0])
		}
	})

	// get last item
	t.Run("Last", func(t *testing.T) {
		n := list.Last()
		if reflect.ValueOf(n.Value()).Int() != values[length-1] {
			t.Errorf("last item should be %d", values[length-1])
		}
	})

	// get next item
	t.Run("Next", func(t *testing.T) {
		n := list.First()
		n = n.Next()
		if reflect.ValueOf(n.Value()).Int() != values[1] {
			t.Errorf("last item should be %d", values[1])
		}
	})

	// get prev item
	t.Run("Prev", func(t *testing.T) {
		n := list.Last()
		n = n.Prev()
		if reflect.ValueOf(n.Value()).Int() != values[length-2] {
			t.Errorf("last item should be %d", values[length-2])
		}
	})

	// get list length
	t.Run("Len", func(t *testing.T) {
		if list.Len() != len(values) {
			t.Errorf("list length should be %d", len(values))
		}
	})

	// add item to end
	t.Run("PushBack", func(t *testing.T) {
		n := list.First()
		for _, v := range values {
			if reflect.ValueOf(n.Value()).Int() != v {
				t.Errorf("first item should be %d", v)
			}
			n = n.Next()
		}
	})

	// add item to top
	t.Run("PushFront", func(t *testing.T) {
		for _, v := range values {
			list.PushFront(v)
		}
		n := list.First()
		for _, v := range values {
			if reflect.ValueOf(n.Value()).Int() != int64(length)+1-v {
				t.Errorf("first item should be %d", v)
			}
			n = n.Next()
		}
	})

	// remove item from front, back
	t.Run("Remove", func(t *testing.T) {
		list := &List{}

		for _, v := range values {
			list.PushBack(v)
		}
		n := list.First()
		n.Remove()
		n = list.First()
		val := n.Value().(int64)
		if val != values[1] {
			t.Errorf("first item should be %d but we have %d", values[1], val)
		}

		n = list.Last()
		n.Remove()
		n = list.Last()
		val = n.Value().(int64)
		if val != values[length-2] {
			t.Errorf("last item should be %d but we have %d", values[length-2], val)
		}

		newlength := list.Len()
		if newlength != length-2 {
			t.Errorf("Length list should be %d but we have %d", length-2, newlength)
		}
	})
}
