//
// Created by aldair on 20/10/20.
//
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int cases, n, resp, size;
    cin >> cases;

    string str;
    while( cases-- ) {
        cin >> str;
        size = str.length();
        n = str[0] - '0';
        resp = (n-1) * 10 + size*(size+1)/2;
        cout << resp << '\n';
    }
    return 0;
}
