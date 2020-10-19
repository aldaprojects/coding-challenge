//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/158/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int n, k;
    scanf("%d%d", &n, &k);

    int a, winners = 0, score = 0;
    for( int i = 0; i < n; i++ ) {
        scanf("%d", &a);
        if( a != 0 ) {
            if( i+1 == k ) {
                score = a;
                winners = k;
            } else if( a >= score ) {
                winners++;
            }
        }
    }

    printf("%d", winners);
    return 0;
}
