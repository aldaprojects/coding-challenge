//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/282/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int n, X = 0;
    cin >> n;

    string statement;
    while( n-- ) {
        cout << statement;
        if( statement[0] == '+' || statement[1] == '+' ) {
            X++;
        } else {
            X--;
        }
    }

    printf("%d", X);
    return 0;
}
