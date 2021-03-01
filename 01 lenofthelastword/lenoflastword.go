package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	bufin := bufio.NewReader(os.Stdin)
	str, err := bufin.ReadString('\n')
	if err != nil {
		return
	}

	//str contains '\n' len 1, so sub 2
	fmt.Println(len(str)-strings.LastIndex(str, " ")-2)
}
