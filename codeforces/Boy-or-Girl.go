//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/236/A
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}
func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func main() {
	defer writer.Flush()

	var username string
	scanf("%s\n", &username)

	characters := make(map[byte]int)

	for i := 0; i < len(username); i++ {
		characters[username[i]]++
	}

	if len(characters)%2 == 0 {
		printf("CHAT WITH HER!\n")
	} else {
		printf("IGNORE HIM!\n")
	}
}
