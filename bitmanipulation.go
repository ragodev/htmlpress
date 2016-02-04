package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hotei/bits"
)

type BitString struct {
	curBit int
	bitSet bits.BitField
}

func Btoi(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func byteString(n int64) string {
	numStr := strconv.FormatInt(n, 2)
	for {
		if len(numStr) == 8 {
			break
		}
		numStr = "0" + numStr
	}
	return numStr
}

func (bits *BitString) AddByte(n int) {
	numStr := byteString(int64(n))
	fmt.Println(numStr)
	for _, c := range numStr {
		if c == 49 {
			bits.bitSet.SetBit(bits.curBit)
		} else {

		}
		bits.curBit += 1
	}
}

func (bits *BitString) AddBit(setIt bool) {
	if setIt {
		bits.bitSet.SetBit(bits.curBit)
	} else {
		bits.bitSet.ClrBit(bits.curBit)
	}
}

func (bits BitString) PrintBytes() {
	sum := float64(0)
	for i := range bits.bitSet.String() {
		multiplier := 7 - math.Mod(float64(i), 8)
		if bits.bitSet.Bit(i) {
			sum += math.Pow(2, multiplier*Btoi(bits.bitSet.Bit(i)))
		}

		if multiplier == 0 {
			fmt.Println(sum)
			sum = 0
		}
	}
}

func (bits BitString) String() string {
	return bits.bitSet.String()
}

func main() {
	// we know the size of this one
	bits := BitString{curBit: 0}
	bits.bitSet.SetMaxBitNdx(8)

	bits.AddByte(4)
	fmt.Println(bits.String())
	bits.PrintBytes()

	bits.AddByte(233)
	fmt.Println(bits.String())
	bits.PrintBytes()

	bits.AddByte(233)
	fmt.Println(bits.String())
	bits.PrintBytes()

	bits.AddByte(233)
	fmt.Println(bits.String())
	bits.PrintBytes()

	bits.AddBit(false)
	fmt.Println(bits.String())
	bits.PrintBytes()

	bits.AddBit(true)
	fmt.Println(bits.String())
	bits.PrintBytes()

}
