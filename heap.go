package binheap

// FirstIndex is the index of the first element. This is helpful
// when calculating parents and children.
const FirstIndex = 1

// OrderingFunc determines how Items are ordered in a binary heap.
type OrderingFunc func(i Item, j Item) bool

// Item represents an element of a binary heap.
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
	if idx, found := b.IndexMap[n]; found {
		parentIdx := idx / 2
		return b.Heap[parentIdx]
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
	childIdx := 2*b.IndexMap[n] + inc
	if 0 < childIdx && childIdx <= b.Size {
		return b.Heap[childIdx]
	}
	return nil
}

func (b *BinaryHeap) swap(first Item, second Item) {
	// we assume the indexes work because this is an unexported function
	firstIdx, _ := b.IndexMap[first]
	secondIdx, _ := b.IndexMap[second]
	b.Heap[firstIdx] = second
	b.Heap[secondIdx] = first
	b.IndexMap[first] = secondIdx
	b.IndexMap[second] = firstIdx
}

func (b *BinaryHeap) heapify(i int) {
	item := b.Heap[i]
	var first Item = item
	var firstIdx int
	// if ordering holds between left and item, set left to first
	if left := b.Left(item); left != nil {
		leftIdx := b.IndexMap[left]
		if b.OrderingHolds(left, item) {
			first = left
			firstIdx = leftIdx
		}
	}
	// if ordering holds between right and first, set right to first
	if right := b.Right(item); right != nil {
		rightIdx := b.IndexMap[right]
		if b.OrderingHolds(right, first) {
			first = right
			firstIdx = rightIdx
		}
	}
	if first != item {
		b.swap(item, first)
		// push the item down to its proper level
		b.heapify(firstIdx)
	}
}
