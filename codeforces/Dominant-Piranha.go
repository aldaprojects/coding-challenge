//
// Created by aldair on 20/10/20.
// https://codeforces.com/problemset/problem/1433/C
//

package main

import (
    "bufio"
    "fmt"
    "os"
)

var (
    reader = bufio.NewReader(os.Stdin)
    writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, a ...interface{}) {
    fmt.Fprintf(writer, f, a...)
}
func scanf(f string, a ...interface{}) {
    fmt.Fscanf(reader, f, a...) 
}

func max(n1, n2 int) int {
    if n2 > n1 {
        return n2
    } else {
        return n1
    }
}

func main() {
    defer writer.Flush()

    var cases, n, piranha int
    scanf("%d\n", &cases)

    for; cases > 0; cases-- {
        scanf("%d\n", &n)
        piranhas := make([]int, 0, n)
        var maxi int
        for i := 0; i < n; i++ {
            scanf("%d", &piranha)
            piranhas = append(piranhas, piranha)
            maxi = max(maxi, piranha)
        }
        scanf("\n")

        var i int
        var dominant bool
        for; i < len(piranhas); i++ {
            if maxi == piranhas[i] {
                if i+1 < len(piranhas) && piranhas[i] > piranhas[i+1] || i-1 >= 0 && piranhas[i] > piranhas[i-1] {
                    dominant = true
                    break
                }
            }
        }
        if dominant {
            printf("%d\n", i+1)
        } else {
            printf("-1\n")
        }
    }
}
