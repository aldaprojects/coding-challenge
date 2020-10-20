//
// Created by aldair on 19/10/20.
// https://codeforces.com/problemset/problem/339/A
//

#include <bits/stdc++.h>

#define fastIO ios_base::sync_with_stdio(0); cin.tie(NULL);

using namespace std;

int main() {
    fastIO

    string str;
    cin >> str;
    
    string nums;
    int size = str.length();
    
    for( int i = 0; i < size; i+=2 ) {
        nums.push_back(str[i]);
    }

    sort(nums.begin(), nums.end());
    for( int i = 0; i < nums.length() - 1; i++ ) {
        printf("%c+", nums[i]);
    }
    printf("%c\n", nums[nums.length()-1]);
    
    return 0;
}
