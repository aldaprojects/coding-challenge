//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/791/A
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

	var weightLimak, weightBob, years int
	scanf("%d%d\n", &weightLimak, &weightBob)
	
	for; weightLimak <= weightBob; {
		weightLimak *= 3
		weightBob *= 2
		years++
	}
	
	printf("%d\n", years)
}
