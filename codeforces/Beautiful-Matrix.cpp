//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/263/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int n1, n2, n3, n4, n5, resp;
    for( int i = 0; i < 5; i++ ) {
        cin >> n1 >> n2 >> n3 >> n4 >> n5;
        resp = abs(2 - i);
        if( n1 || n5 ) {
            resp += 2;
            break;
        } else if ( n2 || n4 ) {
            resp += 1;
            break;
        } else if ( n3 ) {
            break;
        }
    }

    printf("%d\n", resp);
    return 0;
}
