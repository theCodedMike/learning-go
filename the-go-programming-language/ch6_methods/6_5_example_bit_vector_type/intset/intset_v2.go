package intset

import (
	"bytes"
	"fmt"
	"strconv"
)

// An UintSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type UintSet struct {
	words []uint
}

const Bits = 32 << (^uint(0) >> 63)

func getWordAndBit(x uint) (word uint, bit uint) {
	return x / Bits, x % Bits
}

// Has reports whether the set contains the non-negative value x.
func (s *UintSet) Has(x uint) bool {
	word, bit := getWordAndBit(x)
	return word < uint(len(s.words)) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *UintSet) Add(x uint) {
	word, bit := getWordAndBit(x)
	for word >= uint(len(s.words)) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds multiple x to the set
func (s *UintSet) AddAll(elems ...uint) {
	for _, x := range elems {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t.
// 元素要么出现在s集合，要么出现在t集合（并集）
func (s *UintSet) UnionWith(t *UintSet) {
	sLen := len(s.words)
	for i, word := range t.words {
		if i < sLen {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// IntersectWith sets s to intersect of s and t.
// 元素同时出现在s集合与t集合（交集）
func (s *UintSet) IntersectWith(t *UintSet) {
	sLen, tLen := len(s.words), len(t.words)
	if tLen == 0 || sLen == 0 {
		s.Clear()
		return
	}

	for i, word := range t.words {
		if i < sLen {
			s.words[i] &= word
			if i == tLen-1 {
				for j := tLen; j < sLen; j++ {
					s.words[j] = 0
				}
			}
		} else {
			break
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
// 元素出现在s集合但未出现在t集合（差集）
func (s *UintSet) DifferenceWith(t *UintSet) {
	// todo!
	sLen := len(s.words)
	for i, word := range t.words {
		if i < sLen {
			s.words[i] ^= word
		} else {
			break
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
// 元素出现在s集合且未出现在t集合，或元素出现在t集合且未出现在s集合（并差集）
func (s *UintSet) SymmetricDifference(t *UintSet) {
	// todo!
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s UintSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	count := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				count++
				buf.WriteString(strconv.Itoa(64*i + j))
			}
		}
	}

	_, err := fmt.Fprintf(&buf, "}, len: %d", count)
	if err != nil {
		fmt.Printf("Failed to print: %v\n", err)
	}
	return buf.String()
}

// Len returns the number of elements
func (s *UintSet) Len() int {
	num := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if word&(1<<j) != 0 {
				num++
			}
		}
	}
	return num
}

// Remove removes x from the set
func (s *UintSet) Remove(x uint) {
	if s.Has(x) {
		word, bit := getWordAndBit(x)
		s.words[word] ^= 1 << bit
	}
}

// Clear removes all elements from the set
func (s *UintSet) Clear() {
	clear(s.words)
}

// Copy returns a copy of the set
func (s *UintSet) Copy() *UintSet {
	var dst UintSet
	dst.words = make([]uint, len(s.words))
	copy(dst.words, s.words)
	return &dst
}

// Elems return all elems in IntSet
func (s *UintSet) Elems() []uint {
	var elems []uint
	for i, word := range s.words {
		if word != 0 {
			for j := uint(0); j < 64; j++ {
				if word&(1<<j) != 0 {
					elems = append(elems, 64*uint(i)+j)
				}
			}
		}
	}
	return elems
}
