/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only 
store integers within the 32-bit signed integer range: [−231,  231 − 1]. 
For the purpose of this problem, assume that your function returns 0 when 
the reversed integer overflows.
 
*/


namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
        }
        public int Reverse(int x)
        {
            int res = 0;
            int answer = 0;
            while (x != 0)
            {
                res = x % 10;
                x = x / 10;
                if (answer > int.MaxValue / 10 || answer < int.MinValue / 10) return 0;
                answer = answer * 10 + res;
            }
            return answer;
        }

    }
}
