//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/266/B
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

	var coordinate int
	scanf("%d\n", &coordinate)

	var result int

	result += coordinate/5
	coordinate %= 5
	result += coordinate/4
	coordinate %= 4
	result += coordinate/3
	coordinate %= 3
	result += coordinate/2
	coordinate %= 2
	result += coordinate
	printf("%d\n", result)
}
