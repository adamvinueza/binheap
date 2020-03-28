package binheap_test

import (
	"testing"

	. "github.com/adamvinueza/binheap"
	"github.com/stretchr/testify/assert"
)

func TestCompleteTreeNoChange(t *testing.T) {
	items := numerics(3, 4, 5, 6, 7)
	b := NewBinaryHeap(items, func(i Item, j Item) bool {
		return i.Value().(int) <= j.Value().(int)
	})
	assert.True(t, hasCompleteTree(b))
}

func TestCompleteTreeChange(t *testing.T) {
	items := numerics(6, 3, 5, 4, 7)
	b := NewBinaryHeap(items, func(i Item, j Item) bool {
		return i.Value().(int) <= j.Value().(int)
	})
	assert.True(t, hasCompleteTree(b))
}

func TestHeapProperty(t *testing.T) {
	items := numerics(6, 3, 5, 4, 7)
	b := NewBinaryHeap(items, func(i Item, j Item) bool {
		return i.Value().(int) <= j.Value().(int)
	})
	assert.True(t, heapPropertyHolds(b))
}

func TestExtractMax(t *testing.T) {
	items := numerics(4, 6, 3, 7, 5)
	b := NewBinaryHeap(items, func(i Item, j Item) bool {
		return i.Value().(int) >= j.Value().(int)
	})

	for i := 7; i > 2; i-- {
		max := b.ExtractFirst()
		assert.Equal(t, i, max.Value())
		assert.True(t, heapPropertyHolds(b))
		assert.True(t, hasCompleteTree(b))
	}
}

func TestExtractMin(t *testing.T) {
	items := numerics(6, 3, 5, 4, 7)
	b := NewBinaryHeap(items, func(i Item, j Item) bool {
		return i.Value().(int) <= j.Value().(int)
	})

	for i := 3; i < 8; i++ {
		min := b.ExtractFirst()
		assert.Equal(t, i, min.Value())
		assert.True(t, heapPropertyHolds(b))
		assert.True(t, hasCompleteTree(b))
	}
}

type numeric struct {
	value int
}

func (n *numeric) Value() interface{} {
	return n.value
}

func numerics(ii ...int) []Item {
	nn := make([]Item, len(ii))
	for i, v := range ii {
		nn[i] = &numeric{value: v}
	}
	return nn
}

func hasCompleteTree(b *BinaryHeap) bool {
	nodeNotFull := false
	if len(b.Heap) == 0 {
		return true
	}
	queue := make([]Item, 0)
	queue = append(queue, b.Heap[FirstIndex])
	var current Item
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if next := b.Left(current); next != nil {
			// two nodes have only one child
			if nodeNotFull {
				return false
			}
			queue = append(queue, next)
		} else {
			nodeNotFull = true
		}
		if next := b.Right(current); next != nil {
			// node has a right child, but no left child
			if nodeNotFull {
				return false
			}
			queue = append(queue, next)
		} else {
			nodeNotFull = true
		}
	}
	return true
}

func heapPropertyHolds(b *BinaryHeap) bool {
	for _, n := range b.Items() {
		if b.Parent(n) == nil {
			continue
		}
		if !b.OrderingHolds(b.Parent(n), n) {
			return false
		}
	}
	return true
}
