//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/231/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int problems;
    scanf("%d", &problems);

    int a, b, c, numProblems = 0;
    while( problems-- ) {
        scanf("%d%d%d", &a, &b, &c);
        if( a+b+c >= 2 ) {
            numProblems++;
        }
    }

    printf("%d\n", numProblems);
    return 0;
}
