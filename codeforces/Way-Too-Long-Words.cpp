//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/71/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(0); cout.tie(0);

using namespace std;

int main() {
    fastIO

    int cases;
    string word;

    cin >> cases;

    while( cases -- ) {
        cin >> word;
        int wordLength = size(word);
        if( wordLength > 10 ){
            cout << word[0] << wordLength - 2 << word[wordLength - 1] << '\n';
        } else {
            cout << word << '\n';
        }
    }

    return 0;
}