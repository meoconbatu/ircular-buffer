package circular

import "errors"

const testVersion = 4

type Buffer struct {
	array         []byte
	current_index int
	begin_index   int
	length        int
}

func NewBuffer(size int) *Buffer {
	b := Buffer{make([]byte, size), 0, 0, size}
	return &b
}
func (buf *Buffer) ReadByte() (byte, error) {
	b := buf.array[buf.begin_index]
	if b == 0 {
		return 0, errors.New("The buffer is empty!")
	}
	buf.array[buf.begin_index] = 0
	buf.begin_index++
	if buf.begin_index == buf.length {
		buf.begin_index = 0
	}
	return b, nil
}

func (buf *Buffer) WriteByte(c byte) error {
	if buf.array[buf.current_index] != 0 {
		return errors.New("The buffer is full!")
	}
	buf.array[buf.current_index] = c
	buf.current_index++
	if buf.current_index == buf.length {
		buf.current_index = 0
	}
	return nil
}
func (buf *Buffer) Overwrite(c byte) {
	err := buf.WriteByte(c)
	if err != nil {
		buf.array[buf.begin_index] = c
		buf.begin_index++
		if buf.begin_index == buf.length {
			buf.begin_index = 0
		}
	}
}
func (buf *Buffer) Reset() {
	buf.current_index = 0
	buf.begin_index = 0
	buf.array = make([]byte, buf.length)
}
