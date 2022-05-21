package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input01_noPrefix.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	// create buffer
	// b := make([]byte, 100)
	// for {
	// 	//read content to buffer
	// 	readTotal, err := file.Read(b)
	// 	if err != nil {
	// 		// remove eof
	// 		if err != io.EOF{
	// 			fmt.Println(err.Error())
	// 		}
	// 		return
	// 	}

	// 	if string(b[:readTotal]) != "EOF" {
	// 		fmt.Println(string(b[:readTotal]))
	// 	}
	// }

	// use bufio
	scanner := bufio.NewScanner(file)
	// var q int
	var words []string
	for i := 0; scanner.Scan(); i++ {
		if i != 0 {
			words = append(words, scanner.Text())
		}
	}
	// fmt.Println(q)
	// fmt.Println(words)
	noPrefix2(words)
	//noPrefix(words)
}

func noPrefix2(words []string) {
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if len(words[j]) > len(words[i]) {
				if words[i] == words[j][0:len(words[i])] {
					fmt.Println("BAD SET")
					fmt.Println(words[j])
					return
				}
			}
		}
	}
	fmt.Println("GOOD SET")
}

func noPrefix(words []string) {
	// sorting
	//fmt.Println(words)
	sort.Slice(words, func(a, b int) bool { return words[a] < words[b] })
	goodset := true

	for i := 0; i < len(words)-1; i++ {
		if len(words[i+1]) > len(words[i]) {
			if words[i] == words[i+1][0:len(words[i])] {
				fmt.Println("BAD SET")
				fmt.Println(words[i+1])
				break
			}
		}
	}
	if goodset {
		fmt.Println("GOOD SET")
	}

	//fmt.Println(words)
}
