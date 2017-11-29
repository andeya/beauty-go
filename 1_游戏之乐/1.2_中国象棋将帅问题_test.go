package main

import (
	"testing"
)

type SemiByte byte

const (
	// 记忆存储单元的一半，本题为4bit
	HALF_BITS_LENGTH SemiByte = 4
	// 全部bits的mask，即二进制'11111111'
	FULL_MASK SemiByte = 255
	// 右bits的mask，即二进制'00001111'
	R_MASK SemiByte = FULL_MASK >> HALF_BITS_LENGTH
	// 左bits的mask，即二进制'11110000'
	L_MASK SemiByte = R_MASK << HALF_BITS_LENGTH
	// 将帅移动范围的行宽
	GRIDW SemiByte = 3
)

func (s *SemiByte) LGet() SemiByte {
	return (L_MASK & *s) >> HALF_BITS_LENGTH
}

func (s *SemiByte) LSet(n SemiByte) {
	*s = (R_MASK & *s) ^ (n << HALF_BITS_LENGTH)
}

func (s *SemiByte) RGet() SemiByte {
	return R_MASK & *s
}

func (s *SemiByte) RSet(n SemiByte) {
	*s = (L_MASK & *s) ^ n
}

func Test中国象棋将帅问题解法一(t *testing.T) {
	var b SemiByte
	for b.LSet(1); b.LGet() <= GRIDW*GRIDW; b.LSet(b.LGet() + 1) {
		for b.RSet(1); b.RGet() <= GRIDW*GRIDW; b.RSet(b.RGet() + 1) {
			if b.LGet()%GRIDW != b.RGet()%GRIDW {
				t.Logf("A = %d, B = %d\n", b.LGet(), b.RGet())
			}
		}
	}
}
