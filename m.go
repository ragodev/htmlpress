package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/cheggaaa/pb"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func savePats(file string, pats map[string]string) {
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

func loadPats(file string) (pats map[string]string) {
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

	dat, _ := ioutil.ReadFile("doc2.txt")
	html := string(dat)
	// formattedHTML1 := gohtml.Format(string(dat))
	//
	// html := ""
	// for _, line := range strings.Split(formattedHTML1, "\n") {
	// 	html += strings.TrimSpace(line + "\n")
	// }

	dec := make(map[string]string)

	fmt.Println("Running through frequency lists")
	// TODO: START FROM THE BIGGEST WORD

	files := []string{"custom", "5_10_all_rank_interjection.txt", "5_9_all_rank_conjunction.txt", "5_8_all_rank_preposition.txt", "5_7_all_rank_detpro.txt", "5_6_all_rank_determ.txt", "5_5_all_rank_pron.txt", "5_4_all_rank_adverb.txt", "5_3_all_rank_adjective.txt", "5_1_all_rank_noun.txt", "5_2_all_rank_verb.txt"}
	it := 0
	bar := pb.StartNew(6000)
	for _, eachFile := range files {
		dat2, err := ioutil.ReadFile(eachFile)
		check(err)
		for _, line := range strings.Split(string(dat2), "\n") {
			dats := strings.Fields(string(line))
			bar.Increment()
			if len(dats) > 0 {
				// fmt.Println(dats)
				if len(dats[0]) >= 5 {
					it++
					if strings.Contains(html, dats[0]) {
						// i:=(math.Mod(float64(int(it/int(math.Pow(254, 2)))), float64(254)))
						j := byte(math.Mod(float64(int(it/254)), float64(254)))
						k := byte(math.Mod(float64(it), float64(254)))
						stringToEncode := dats[0]
						encodedString := string([]byte{171, j, k})
						dec[encodedString] = stringToEncode
						html = strings.Replace(html, stringToEncode, encodedString, -1)
					}
					if strings.Contains(html, strings.Title(dats[0])) {
						// i:=(math.Mod(float64(int(it/int(math.Pow(254, 2)))), float64(254)))
						j := byte(math.Mod(float64(int(it/254)), float64(254)))
						k := byte(math.Mod(float64(it), float64(254)))
						stringToEncode := strings.Title(dats[0])
						encodedString := string([]byte{173, j, k})
						dec[encodedString] = stringToEncode
						html = strings.Replace(html, stringToEncode, encodedString, -1)
					}
					if strings.Contains(html, strings.ToUpper(dats[0])) {
						// i:=(math.Mod(float64(int(it/int(math.Pow(254, 2)))), float64(254)))
						j := byte(math.Mod(float64(int(it/254)), float64(254)))
						k := byte(math.Mod(float64(it), float64(254)))
						stringToEncode := strings.ToUpper(dats[0])
						encodedString := string([]byte{174, j, k})
						dec[encodedString] = stringToEncode
						html = strings.Replace(html, stringToEncode, encodedString, -1)
					}

				}
			}
		}
	}
	bar.FinishPrint("Finished parsing subsets")
	savePats("dec", dec)
	fmt.Println(html)

	// dec := loadPats("dec")
	//
	// for key := range dec {
	// 	if strings.Contains(html, dec[key]) {
	// 		stringToEncode := dec[key]
	// 		html = strings.Replace(html, stringToEncode, key, -1)
	// 	}
	// 	if strings.Contains(html, strings.Title(dec[key])) {
	// 		stringToEncode := strings.Title(dec[key])
	// 		html = strings.Replace(html, stringToEncode, key, -1)
	// 	}
	// 	if strings.Contains(html, strings.ToUpper(key)) {
	// 		stringToEncode := strings.ToUpper(dec[key])
	// 		html = strings.Replace(html, stringToEncode, key, -1)
	// 	}
	// }

	fmt.Println(len(dec))
	d1 := []byte(html)
	err := ioutil.WriteFile("doc2.enc", d1, 0644)
	check(err)

	dat3, _ := ioutil.ReadFile("doc2.enc")
	newhtml := string(dat3)
	for key := range dec {
		newhtml = strings.Replace(newhtml, key, dec[key], -1)
	}
	d2 := []byte(newhtml)
	err = ioutil.WriteFile("doc2.dec", d2, 0644)
	check(err)
}
