//
// Created by aldair on 21/10/20.
// https://codeforces.com/problemset/problem/734/A
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


	var games, anton, danik int
	scanf("%d\n", &games)

	var str string
	scanf("%s\n", &str)
	for i := 0; i < games; i++ {
		if str[i] == 'A' {
			anton++
		} else {
			danik++
		}
	}

	if anton > danik {
		printf("Anton\n")
	} else if anton < danik {
		printf("Danik\n")
	} else {
		printf("Friendship\n")
	}
}
