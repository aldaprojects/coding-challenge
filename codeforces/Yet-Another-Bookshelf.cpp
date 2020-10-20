//
// Created by aldair on 20/10/20.
//
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    int cases, size;
    cin >> cases;
    int n, z1, z2;


    bool b = false;
    while( cases-- ) {
        cin >> size;
        z1 = 0, z2 = 0;
        while( size-- ) {
            cin >> n;
            if( n ) {
                b = true;
                z2 = 0;
            }
            if( b ) {
                if( n < 1 ) {
                    z1++;
                    z2++;
                }
            }
        }
        cout << z1-z2 << endl;
    }

    return 0;
}
