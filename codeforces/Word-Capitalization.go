//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/281/A
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
	
	var word string
	scanf("%s\n", &word)

	word = strings.Title(word)
	printf("%s\n", word)
}
