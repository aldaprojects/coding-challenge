using System;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
    
        }
        public bool IsPalindrome(int x)
        {

            char[] charArray = x.ToString().ToCharArray();
            Array.Reverse(charArray);

            string palindromo = new string(charArray);

            return palindromo == x.ToString();
        }
    }
}

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
 */
