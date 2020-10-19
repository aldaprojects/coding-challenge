//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/112/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    string str1, str2;
    cin >> str1 >> str2;

    int cont = 0, size = str1.length();
    for( int i = 0; i < size; i++ ) {
        if( str1[i] < 97 ) {
            str1[i] += 32;
        }
        if( str2[i] < 97 ) {
            str2[i] += 32;
        }

        if( str1[i] > str2[i] ) {
            cont = 1;
            break;
        } else if( str1[i] < str2[i] ) {
            cont = -1;
            break;
        }
    }

    printf("%d", cont);
    return 0;
}


