//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/116/A
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

	var stops int
	scanf("%d\n", &stops)

	var exit, enter, capacity, aux int
	for; stops > 0; stops-- {
		scanf("%d%d\n", &exit, &enter)

		aux += -exit + enter
		if aux > capacity {
			capacity = aux
		}
	}
	
	printf("%d\n", capacity)
}
