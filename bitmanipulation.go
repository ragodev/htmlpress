package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hotei/bits"
	"github.com/yosssi/gohtml"
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

func Ctoi(n int) bool {
	if n == 49 {
		return true
	}
	return false
}

func bitStringSum(s string) int {
	sum := float64(0)
	numBits := float64(len(s))
	for i, c := range s {
		multiplier := numBits - 1 - math.Mod(float64(i), numBits)
		if Ctoi(int(c)) {
			sum += math.Pow(2, multiplier)
		}
	}
	return int(sum)
}

func byteString(n int64, numBits int) string {
	numStr := strconv.FormatInt(n, 2)
	for {
		if len(numStr) == numBits {
			break
		}
		numStr = "0" + numStr
	}
	return numStr
}

func (bits *BitString) AddBits(n int, numBits int) {
	numStr := byteString(int64(n), numBits)
	for _, c := range numStr {
		if c == 49 {
			bits.bitSet.SetBit(bits.curBit)
		} else {

		}
		bits.curBit += 1
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

func writeBytes(s string) {
	var bs []byte
	i := 0
	for {
		// fmt.Println(s[i : i+8])
		bs = append(bs, byte(bitStringSum(s[i:i+8])))
		i += 8
		if i >= len(s) {
			break
		}
	}
	// fmt.Println(bs)
	f, _ := os.Create("dat")
	defer f.Close()
	f.Write(bs)
}

func readBytes() string {
	dats, _ := ioutil.ReadFile("dat")
	// fmt.Println(dats)
	bsString := ""
	for _, dat := range dats {
		bsString += byteString(int64(dat), 8)

	}
	return bsString
}

func main() {
	// formattedHTML := gohtml.Format(`<html>hello<p>I said Hello</p><div>What is going on and hello there</div></html>`)
	dat, _ := ioutil.ReadFile("html")
	formattedHTML1 := gohtml.Format(string(dat))

	formattedHTML := ""
	for _, line := range strings.Split(formattedHTML1, "\n") {
		formattedHTML += strings.TrimSpace(line + "\n")
	}

	m := make(map[string]int)
	decode := make(map[int]string)
	m["tally123"] = 0

	// ENCODING
	fmt.Println("encoding")
	bits := BitString{curBit: 0}
	bits.bitSet.SetMaxBitNdx(8)
	curWord := ""
	for _, c := range strings.TrimSpace(formattedHTML) {
		strC := string(c)
		if len(strings.TrimSpace(strC)) == 0 || strings.Contains("!#$%&'()*+,-.:;=?@[/\\]^_`{|}~><' \n\t\b", strC) || strings.Contains(`""`, strC) {
			if len(curWord) > 2 {
				word := strings.ToLower(curWord)
				fmt.Println("|" + strings.TrimSpace(curWord) + "|")
				if _, ok := m[word]; ok {
					// has key
				} else {
					m[word] = m["tally123"]
					decode[m["tally123"]] = word
					m["tally123"]++
				}
				bits.AddBits(0, 1)
				if curWord == strings.Title(word) {
					bits.AddBits(1, 1)
					bits.AddBits(0, 1)
				} else if curWord == strings.ToUpper(word) {
					bits.AddBits(0, 1)
					bits.AddBits(1, 1)
				} else {
					bits.AddBits(0, 1)
					bits.AddBits(0, 1)
				}
				bits.AddBits(m[word], 9)
				curWord = ""
			} else if len(curWord) > 0 {
				// write the character bytes that aren't quite enough to be a word
				for _, d := range curWord {
					bits.AddBits(1, 1)
					bits.AddBits(int(d), 8)
				}
				curWord = ""
			}
			// write the current character byte
			if int(c) == 32 {
				// space
				bits.AddBits(0, 1)
				bits.AddBits(1, 1)
				bits.AddBits(1, 1)
			} else {
				bits.AddBits(1, 1)
				bits.AddBits(int(c), 8)
			}
		} else {
			curWord += strC
		}

	}
	fmt.Println(m)

	fmt.Println("writing bytes")
	fmt.Println(len(dat))
	fmt.Println(len(bits.String()))
	writeBytes(bits.String())
	fmt.Println("reading bytes")
	bsString := readBytes()
	fmt.Println(len(bsString))
	i := 0
	for {
		if bsString[i] == 49 {
			i++
			nextByte := bitStringSum(bsString[i : i+8])
			fmt.Print(string(nextByte))
			i += 8
		} else {
			i++
			firstBit := bitStringSum(string(bsString[i]))
			i++
			secondBit := bitStringSum(string(bsString[i]))
			i++
			decodedWord := ""
			if firstBit == 1 && secondBit == 0 {
				decodedWord = decode[bitStringSum(bsString[i:i+9])]
				decodedWord = strings.Title(decodedWord)
				i += 9
			} else if firstBit == 0 && secondBit == 1 {
				decodedWord = decode[bitStringSum(bsString[i:i+9])]
				decodedWord = strings.ToUpper(decodedWord)
				i += 9
			} else if firstBit == 1 && secondBit == 1 {
				decodedWord = " "
			} else {
				decodedWord = decode[bitStringSum(bsString[i:i+9])]
				i += 9
			}
			fmt.Print(decodedWord)
		}
		if i >= len(bsString)-8 {
			break
		}
	}

}
