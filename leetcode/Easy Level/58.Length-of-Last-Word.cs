/*
Given a string s consists of upper/lower-case alphabets 
and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/

public class Solution {
    public int LengthOfLastWord(string s) {
        if(string.IsNullOrEmpty(s)) return 0;
        
        int whiteSpaces = 0;
        for(int i = s.Length-1; i>=0; i--){
            if(char.IsWhiteSpace(s[i])) whiteSpaces++;
            else break;
        }
        
        s = s.Substring(0, s.Length - whiteSpaces);
        
        int size = 0;
        for(int i = s.Length -1; i>=0; i--){
            if(!char.IsWhiteSpace(s[i])) size++;
            else break;
        }
        
        return size;
    }
}
