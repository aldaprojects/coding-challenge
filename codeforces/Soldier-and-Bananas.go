//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/546/A
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

	var cost, dollars, bananas int
	scanf("%d%d%d", &cost, &dollars, &bananas)
	
	borrow := getTotal(cost, bananas) - dollars
	if borrow > 0 {
		printf("%d\n", borrow)
	} else {
		printf("0\n")
	}

}

func getTotal(cost, bananas int) int {
	var total int
	for i := 1; i <= bananas; i++ {
		total += cost * i
	}
	return  total
}