//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/41/A
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

func reverse(s string) string {
	runes := make([]rune, 0, len(s))
	for i := len(s)-1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}
	return string(runes)
}

func main() {
	defer writer.Flush()

	var str1, str2 string
	scanf("%s\n%s", &str1, &str2)

	if reverse(str1) == str2 {
		printf("YES\n")
	} else {
		printf("NO\n")
	}
}
