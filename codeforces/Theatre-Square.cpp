//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/1/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(0); cout.tie(0);

using namespace std;

int main() {
    fastIO

    long long int n, m, a, width, height;
    scanf("%lli%lli%lli", &n, &m, &a);


    if( n%a == 0 ) {
        width = n/a;
    } else {
        width = n/a + 1;
    }

    if( m%a == 0 ) {
        height = m/a;
    } else {
        height = m/a + 1;
    }

    printf("%lli\n", width * height);
    return 0;
}
