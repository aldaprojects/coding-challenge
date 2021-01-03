//
// Created by aldair on 12/21/2020.
// https://www.hackerrank.com/challenges/counting-valleys/problem?h_
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

	
}
