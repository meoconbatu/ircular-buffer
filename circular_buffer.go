package circular

import "errors"

const testVersion = 4

type Buffer struct {
	array  []byte
	tail   int
	head   int
	length int
}

func (buf *Buffer) IsEmpty() bool {
	b := buf.array[buf.head]
	if b == 0 {
		return true
	}
	return false
}
func (buf *Buffer) IsFull() bool {
	if buf.array[buf.tail] != 0 {
		return true
	}
	return false
}
func NewBuffer(size int) *Buffer {
	b := Buffer{make([]byte, size), 0, 0, size}
	return &b
}
func (buf *Buffer) ReadByte() (byte, error) {
	if buf.IsEmpty() {
		return 0, errors.New("The buffer is empty!")
	}
	b := buf.array[buf.head]
	buf.array[buf.head] = 0
	buf.head++
	if buf.head == buf.length {
		buf.head = 0
	}
	return b, nil
}

func (buf *Buffer) WriteByte(c byte) error {
	if buf.IsFull() {
		return errors.New("The buffer is full!")
	}
	buf.array[buf.tail] = c
	buf.tail++
	if buf.tail == buf.length {
		buf.tail = 0
	}
	return nil
}
func (buf *Buffer) Overwrite(c byte) {
	err := buf.WriteByte(c)
	if err != nil {
		buf.array[buf.head] = c
		buf.head++
		if buf.head == buf.length {
			buf.head = 0
		}
	}
}
func (buf *Buffer) Reset() {
	buf.head, buf.tail = 0, 0
	buf.array = make([]byte, buf.length)
}
