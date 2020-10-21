//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/977/A
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

	var n, subtractions int
	scanf("%d%d\n", &n, &subtractions)

	for i := 0; i < subtractions; i++ {
		if n%10 != 0 {
			n--
		} else {
			n /= 10
		}
	}
	
	printf("%d\n", n)
	
}
