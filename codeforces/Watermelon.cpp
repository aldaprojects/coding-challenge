//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/4/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(0); cout.tie(0);

using namespace std;

int main () {
    fastIO

    int n;
    cin >> n;

    if ( n > 2 && n%2 == 0 ) {
        cout << "YES\n";
    } else {
        cout << "NO\n";
    }

    return 0;
}