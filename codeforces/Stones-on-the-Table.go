//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/266/A
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

	var n int
	scanf("%d\n", &n)

	var stones string
	scanf("%s\n", &stones)

	var count int
	for i := 0; i < len(stones)-1; i++ {
		if stones[i] == stones[i+1] {
			count++
		}
	}
	
	printf("%d\n", count)
}
