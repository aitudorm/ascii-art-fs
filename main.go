package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		fmt.Println("Please, enter only 2 arguments \"[STRING]\" \"[TYPE OF BANNER]\" \nP.S String word should be quoted, otherwise keep ascii unicode i.e., `\\\\n` `\\\\t` `\\\\` so on. ")
		return
	} else if len(arg) > 2 {
		fmt.Println("Please, enter only 2 arguments \"[STRING]\" \"[TYPE OF BANNER]\" \nP.S String word should be quoted, otherwise keep ascii unicode i.e., `\\\\n` `\\\\t` `\\\\` so on. ")
		return
	}
	var content []byte
	var err error
	str := strings.ReplaceAll(arg[0], "\\n", "\n")
	banner := arg[1]
	switch banner {
	case "standard":
		content, err = ioutil.ReadFile("banner/standard.txt")
		if err != nil {
			fmt.Println("Cannot Read standard.txt file")
		}
	case "shadow":
		content, err = ioutil.ReadFile("banner/shadow.txt")
		if err != nil {
			fmt.Println("Cannot Read shadaw.txt file")
		}
	case "thinkertoy":
		content, err = ioutil.ReadFile("banner/thinkertoy.txt")
		if err != nil {
			fmt.Println("Cannot Read thinkertoy.txt file")
		}
	default:
		content, err = ioutil.ReadFile("banner/shadow.txt")
		if err != nil {
			fmt.Println("Cannot Read thinkertoy.txt file")
		}
	}

	if !checkStdFile(string(content)) {
		fmt.Println("standard.txt file is corrupted")
		return
	}

	if len(str) == 0 {
		return
	}

	for _, l := range str {
		if l < 32 || l > 126 {
			if l == 10 {
				continue
			}
			fmt.Println("Please enter only enter characters from 32 to 126")
			return
		}
	}

	isThereNewLine, _ := checkNewline(str)
	words1 := strings.Split(str, "\n")
	if isThereNewLine {
		if onlyNewlines(str) {
			words := []string{}

			for i := 1; i < len(words1); i++ {
				words = append(words, words1[i])
			}
			for i := 0; i < len(words); i++ {
				fmt.Println()
			}
		} else {
			for i := 0; i < len(words1); i++ {
				if words1[i] == "" {
					fmt.Println()
					continue
				} else {
					printWord(string(content), words1[i], banner)
				}
			}
		}
	} else {
		printWord(string(content), str, banner)
	}
}

func printWord(content string, str string, banner string) {
	strArr := [8]string{}
	var fontTxt []string
	if banner == "thinkertoy" {
		fontTxt = strings.Split(string(content), "\r\n")
	} else {
		fontTxt = strings.Split(string(content), "\n")

	}
	for _, l := range str {
		pos := int(l)*9 - 287
		if l == 10 {
			continue
		}
		for i := 0; i < 8; i++ {
			strArr[i] += fontTxt[i+pos]
		}
	}
	for i := range strArr {
		fmt.Println(strArr[i])
	}
}

func checkNewline(str string) (bool, int) {
	flag := false
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			flag = true
			count++
		}
	}
	return flag, count
}

func checkStdFile(content string) bool {
	hasher := sha256.New()
	s, err := ioutil.ReadFile("banner/standard.txt")
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	l := hasher.Sum(nil)

	hash_std := []byte{195, 236, 117, 132, 251, 126, 207, 189, 115, 158, 107, 63, 111, 99, 253, 31, 229, 87, 210, 174, 62, 36, 248, 112, 115, 13, 156, 248, 178, 85, 158, 148}

	return string(hash_std) == string(l)
}

func onlyNewlines(s string) bool {
	for _, l := range s {
		if l != '\n' {
			return false
		}
	}
	return true
}
