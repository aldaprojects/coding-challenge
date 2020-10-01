/*
The count-and-say sequence is the sequence of 
integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of 
the count-and-say sequence. You can do so recursively, 
in other words from the previous member read off the digits, 
counting the number of digits in groups of the same digit.

Note: Each term of the sequence of integers will be represented as a string.
*/
class Solution {
public:
    string countAndSay(int n) {
        string seq = "1";
        for(int i = 1; i < n; i++){
            int cont = 0;
            char temp = seq[0];
            string term = "";
            for(int j = 0; j < seq.size(); j++){
                if(seq[j] == temp) cont++;
                else {
                    temp = seq[j];
                    term += to_string(cont) + seq[j-1];
                    cont = 0;
                    j--;
                }
            }
            term += to_string(cont) + seq[seq.size()-1];
            seq = term;
        }
        return seq;
    }
};
