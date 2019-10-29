package main

// Buffer shared resource
type Buffer struct {
	data    []int
	MaxSize int
	extract int
	insert  int

	size int
}

// Size gives the number of elements in the buffer
func (b *Buffer) Size() int {
	return b.size
}

// Insert inserts the given item into the buffer
func (b *Buffer) Insert(item int) {
	b.data[b.insert] = item

	b.insert = (b.insert + 1) % (b.MaxSize)
	b.size++
}

// Extract returns the first item ready to be extracted
func (b *Buffer) Extract() int {
	out := b.data[b.extract]
	b.extract = (b.extract + 1) % (b.MaxSize)
	b.size--

	return out
}
