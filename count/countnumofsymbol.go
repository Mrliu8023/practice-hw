package main

import (
	"fmt"
	"io"
	"os"
)

//
//题目描述
//写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写。
//
//输入描述:
//第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母。
//
//输出描述:
//输出输入字符串中含有该字符的个数。

func main() {
	letterMap := make(map[byte]int)
	letter := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(letter)
		if err == io.EOF || letter[0] == '\n' {
			break
		}
		if err != nil {

			return
		}

		letterMap[convert(letter[0])] += 1
	}

	_, err := os.Stdin.Read(letter)
	if err != nil {
		return
	}

	fmt.Println(letterMap[convert(letter[0])])
}

func convert(b byte) byte {
	// 'a' == 97 'A' == 65
	if b >= 97 {
		b = b - 32
	}
	return b
}
