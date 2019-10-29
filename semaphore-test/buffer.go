package main

// Buffer shared resource
type Buffer struct {
	data    []int
	MaxSize int
	extract int
	insert  int
}

// Size gives the number of elements in the buffer
func (b *Buffer) Size() int {
	return len(b.data)
}

// Insert inserts the given item into the buffer
func (b *Buffer) Insert(item int) {
	b.data[b.insert] = item

	b.insert = (b.insert + 1) % (b.MaxSize)
}

// Extract returns the first item ready to be extracted
func (b *Buffer) Extract() int {
	out := b.data[b.extract]
	b.extract = (b.extract + 1) % (b.MaxSize)

	return out
}
