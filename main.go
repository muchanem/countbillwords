package main

import ( 
	"fmt"
	"time"
	"bufio"
	"log"
	"os"
)

func billtoslice(path string)([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func wordpercent(billtext []string) float64 {
	for i, w := range billtext {
		if w == "SEC." {
			if billtext[i+1] == (os.Args[1] + ".") {
				per := (float64(i)/ float64(len(billtext))* 100)
				return per
			}
		}
	}
	return 0.00
}

func timeremaining(percent float64) float64 {
	mins := time.Since(time.Date(2021,3,4,20,21,3,0,time.UTC)).Minutes()
	return (((100.00/percent) * mins) - mins)
}

func main() {
	billtext , err := billtoslice("billtext.txt")
	if err != nil {
        log.Fatalf("readWords: %s", err)
	}
	wordper := wordpercent(billtext)
	fmt.Println(time.Since(time.Date(2021,3,4,20,21,3,0,time.UTC)).Hours())
	fmt.Println(wordper)
	fmt.Println(timeremaining(wordper)/60)

	//fmt.Println(timeremaining())
}