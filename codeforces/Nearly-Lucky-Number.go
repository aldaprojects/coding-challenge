//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/110/A
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

	var n string
	scanf("%s\n", &n)

	var luckyDigits int
	for i := 0; i < len(n); i++ {
		if n[i] == '4' || n[i] == '7' {
			luckyDigits++
		}
	}

	if luckyDigits == 4 || luckyDigits == 7 {
		printf("YES\n")
	} else {
		printf("NO\n")
	}
}
