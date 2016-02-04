package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/yosssi/gohtml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func savePats(file string, pats map[string]int) {
	f, err := os.Create(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(pats); err != nil {
		panic("cant encode")
	}
}

func loadPats(file string) (pats map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err := enc.Decode(&pats); err != nil {
		panic("cant decode")
	}

	return pats
}

func main() {
	nextFileNum := "4"
	f, _ := os.Create("dat" + nextFileNum)

	defer f.Close()

	dat, err := ioutil.ReadFile("doc" + nextFileNum + ".txt")
	check(err)
	formattedHTML := gohtml.Format(string(dat))
	m := make(map[string]int)
	m["tally123"] = 0
	// strInt, _ := strconv.Atoi(nextFileNum)
	// m := loadPats("dat" + strconv.Itoa(strInt-1) + ".encoding")
	curWord := ""
	for _, c := range strings.TrimSpace(formattedHTML) {
		strC := string(c)
		if len(strings.TrimSpace(strC)) == 0 || strings.Contains("!#$%&'()*+,-.:;=?@[/\\]^_`{|}~><' \n\t\b", strC) || strings.Contains(`""`, strC) {
			if len(curWord) > 0 {
				word := strings.ToLower(curWord)
				fmt.Println("|" + strings.TrimSpace(curWord) + "|")
				if _, ok := m[word]; ok {
					// has key
				} else {
					m[word] = m["tally123"]
					m["tally123"] += 1
				}
				i := byte(0)
				if curWord == strings.Title(word) {
					i = byte(1)
				} else if curWord == strings.ToUpper(word) {
					i = byte(2)
				}
				j := byte(m[word] / 254)
				k := byte(math.Mod(float64(m[word]), 254))
				fmt.Println(i)
				d2 := []byte{j, k}
				f.Write(d2)
				curWord = ""
			} else {
				d2 := []byte{byte(c)}
				f.Write(d2)
			}
		} else {
			curWord += strC
		}

	}
	fmt.Println(curWord)

	savePats("dat"+nextFileNum+".encoding", m)
	fmt.Println(m)
	fmt.Println(m["tally123"])
}
