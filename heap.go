package binheap

// FirstIndex is the index of the first element. This is helpful
// when calculating parents and children.
const FirstIndex = 1

// OrderingFunc determines how items are ordered in a binary heap.
type OrderingFunc func(i Item, j Item) bool

// Items go into a binary heap.
//
// Value() retrieves the value that can be compared in the binary heap's
// ordering function. It is presumed that the ordering function knows which
// types to expect.
type Item interface {
	Value() interface{}
}

// BinaryHeap implements a binary heap.
type BinaryHeap struct {
	Heap          []Item
	IndexMap      map[Item]int
	OrderingHolds OrderingFunc
	Length        int
	Size          int
}

// NewBinaryHeap returns a BinaryHeap containing the specified items, inserted
// using the specified OrderingFunc.
func NewBinaryHeap(items []Item, fn OrderingFunc) *BinaryHeap {
	b := BinaryHeap{
		Heap:          append([]Item{nil}, items...),
		OrderingHolds: fn,
		IndexMap:      make(map[Item]int),
	}
	for i := 1; i < len(b.Heap); i++ {
		item := b.Heap[i]
		b.IndexMap[item] = i
	}
	b.build()
	return &b
}

// Left returns the left node in this BinaryHeap.
func (b *BinaryHeap) Left(n Item) Item {
	return b.child(n, 0) // 2i
}

// Right returns the right node in this BinaryHeap.
func (b *BinaryHeap) Right(n Item) Item {
	return b.child(n, 1) // 2i + 1
}

// Parent returns the parent of the specified item.
func (b *BinaryHeap) Parent(n Item) Item {
	if idx, found := b.index(n); found {
		return b.Heap[idx/2]
	}
	return nil
}

// ExtractFirst removes and returns the first Item from this BinaryHeap.
func (b *BinaryHeap) ExtractFirst() Item {
	first := b.Heap[FirstIndex]
	b.swap(first, b.Heap[b.Size])
	b.Size--
	b.heapify(FirstIndex)
	return first
}

// Items returns this BinaryHeaps items, in order of its internal Heap slice.
func (b *BinaryHeap) Items() []Item {
	items := make([]Item, 0)
	for i := 1; i <= b.Size; i++ {
		items = append(items, b.Heap[i])
	}
	return items
}

// build takes a slice of heap Items and reorders them into a binary heap.
func (b *BinaryHeap) build() {
	b.Size = len(b.Heap) - 1
	b.Length = len(b.Heap) - 1
	for i := b.Length / 2; i > 0; i-- {
		b.heapify(i)
	}
}

func (b *BinaryHeap) child(n Item, inc int) Item {
	if idx, found := b.index(n); found {
		childIdx := idx*2 + inc
		if childIdx < len(b.Heap) {
			return b.Heap[childIdx]
		}
	}
	return nil
}

// index returns the index of the item in this BinaryHeap's internal slice.
func (b *BinaryHeap) index(n Item) (int, bool) {
	idx, found := b.IndexMap[n]
	return idx, found
}

func (b *BinaryHeap) swap(first Item, second Item) {
	// we assume the indexes work because this is an unexported function
	firstIdx, _ := b.index(first)
	secondIdx, _ := b.index(second)
	b.Heap[firstIdx] = second
	b.Heap[secondIdx] = first
	b.IndexMap[first] = secondIdx
	b.IndexMap[second] = firstIdx
}

// heapify rebuilds the internal heap slice.
func (b *BinaryHeap) heapify(i int) {
	item := b.Heap[i]
	left := b.Left(item)
	right := b.Right(item)
	var largest Item = item
	var largestIdx int
	if idx, found := b.index(left); found {
		if idx <= b.Size && b.OrderingHolds(left, item) {
			largest = left
			largestIdx = idx
		}
	}
	if idx, found := b.index(right); found {
		if idx <= b.Size && b.OrderingHolds(right, largest) {
			largest = right
			largestIdx = idx
		}
	}
	if largest != item {
		b.swap(item, largest)
		b.heapify(largestIdx)
	}
}
