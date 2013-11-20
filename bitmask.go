package lape

import (
        `bytes`
        //`fmt`
)

type Bitmask uint64

// Returns true if all bitmask bits are clear.
func (b Bitmask) IsEmpty() bool {
	return b == Bitmask(0)
}

// Returns true if all bitmask bits are set.
func (b Bitmask) IsFull() bool {
	return b == 0xFFFFFFFFFFFFFFFF
}

// Returns true if a bit at given position is set.
func (b Bitmask) IsSet(position int) bool {
	return b & (1 << uint(position)) != Bitmask(0)
}

// Returns true if a bit at given position is clear.
func (b Bitmask) IsClear(position int) bool {
	return !b.IsSet(position)
}

func (b Bitmask) ToString() string {
	buffer := bytes.NewBufferString("  a b c d e f g h\n")
	for row := 7; row >= 0; row-- {
		buffer.WriteByte('1' + byte(row))
		for col := 0; col <= 7; col++ {
			position := row << 3 + col
			buffer.WriteByte(' ')
			if b.IsSet(position) {
				buffer.WriteString("\u2022") // Set
			} else {
				buffer.WriteString("\u22C5") // Clear
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

// Sets a bit at given position.
func (b *Bitmask) Set(position int) {
	*b |= 1 << uint(position)
}

// Clears a bit at given position.
func (b *Bitmask) Clear(position int) {
	*b &= ^(1 << uint(position))
}

// Combines two bitmasks using bitwise OR operator.
func (b *Bitmask) Combine(bitmask Bitmask) {
	*b |= bitmask
}

// Intersects two bitmasks using bitwise AND operator.
func (b *Bitmask) Intersect(bitmask Bitmask) {
	*b &= bitmask
}

// Mulitplies two bitmasks.
func (b *Bitmask) Multiply(bitmask Bitmask) {
	*b *= bitmask
}

// Excludes bits of one bitmask from another using bitwise XOR operator.
func (b *Bitmask) Exclude(bitmask Bitmask) {
	*b ^= (bitmask & *b)
}

// Finds out row number of bit position.
func Row(position int) int {
	return position / 8 // position >> 3
}

// Finds out column number of bit position.
func Column(position int) int {
	return position % 8 // position & 7
}

// Finds out bit position for given row and column.
func Index(row, column int) int {
	return (row << 3) + column
}

func Abs(i int) int {
        if i >= 0 {
                return i
        } else {
                return -i
        }
}
