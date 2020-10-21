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

	var n, t int
	scanf("%d%d\n", &n, &t)

	var str []byte
	scanf("%s\n", &str)

	for; t > 0; t-- {
		for i := 0; i < len(str)-1; i++ {
			if str[i] == 'B' && str[i+1] == 'G' {
				str[i], str[i+1] = str[i+1], str[i]
				i++
			}
		}
	}

	printf("%s\n", str)
}
