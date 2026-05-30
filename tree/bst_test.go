package tree

import (
	"reflect"
	"testing"
)

func TestBSTInsertAndInOrder(t *testing.T) {
	b := NewBST(5, 3, 7, 1, 4, 6, 8)
	want := []int{1, 3, 4, 5, 6, 7, 8}
	if got := b.InOrder(); !reflect.DeepEqual(got, want) {
		t.Errorf("InOrder() = %v, want %v", got, want)
	}
}

func TestBSTSearch(t *testing.T) {
	b := NewBST(5, 3, 7, 1, 4)

	cases := []struct {
		val  int
		want bool
	}{
		{4, true},
		{7, true},
		{1, true},
		{99, false},
		{0, false},
	}
	for _, tt := range cases {
		if got := b.Search(tt.val); got != tt.want {
			t.Errorf("Search(%d) = %v, want %v", tt.val, got, tt.want)
		}
	}
}

func TestBSTDelete(t *testing.T) {
	t.Run("leaf", func(t *testing.T) {
		b := NewBST(5, 3, 7, 1)
		if !b.Delete(1) {
			t.Fatal("Delete(1) = false, want true")
		}
		want := []int{3, 5, 7}
		if got := b.InOrder(); !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("one child", func(t *testing.T) {
		b := NewBST(5, 3, 7, 1, 4)
		b.Delete(3)
		want := []int{1, 4, 5, 7}
		if got := b.InOrder(); !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("two children", func(t *testing.T) {
		b := NewBST(5, 3, 7, 1, 4, 6, 8)
		b.Delete(5)
		want := []int{1, 3, 4, 6, 7, 8}
		if got := b.InOrder(); !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder() = %v, want %v", got, want)
		}
	})

	t.Run("missing", func(t *testing.T) {
		b := NewBST(5, 3)
		if b.Delete(99) {
			t.Error("Delete(99) = true, want false")
		}
	})
}

func TestBSTEmpty(t *testing.T) {
	b := &BST{}
	if b.Search(1) {
		t.Error("Search on empty tree should be false")
	}
	if got := b.InOrder(); len(got) != 0 {
		t.Errorf("InOrder() = %v, want []", got)
	}
}
