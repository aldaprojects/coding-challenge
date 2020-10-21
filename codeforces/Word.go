//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/59/A
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var word []byte
	scanf("%s\n", &word)

	var upper, lower int
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			upper++
		} else {
			lower++
		}
	}

	if upper > lower {
		printf("%s\n", strings.ToUpper(string(word)))
	} else {
		printf("%s\n", strings.ToLower(string(word)))
	}
}
