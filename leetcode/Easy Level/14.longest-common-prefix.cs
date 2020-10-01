using System;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {

        }
        public string LongestCommonPrefix(string[] strs)
        {
            string prefix = "";

            if (strs.Length == 0) return prefix;
            if (strs.Length == 1) return strs[1];

            for (int i = 0; i < strs[0].Length; i++)
            {
                for (int j = 1; j < strs.Length; j++)
                {
                    if (i >= strs[j].Length || strs[j][i] != strs[0][i])
                    {
                        return prefix;
                    }
                }
                prefix = prefix + strs[0][i];
            }
            return prefix;
        }

    }
}

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z. 
     
*/
