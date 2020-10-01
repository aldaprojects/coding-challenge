/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/


using System.Collections.Generic;

namespace solution {
    class Program {

        static void Main(string[] args) {

            Dictionary<char, char> map = new Dictionary<char,char>();

            map.Add('{', '}');
            map.Add('(', ')');
            map.Add('[', ']');

            string prueba = "{[]}";

            System.Console.WriteLine( validar(prueba, ref map));
            
        }

        static bool validar(string s, ref Dictionary<char, char> map){

            if(s == "") return true;
            bool valido = false;
            
            for(int i = 0; i < s.Length; i++) {
                valido = false;
                string expresion = "";
                if( map.ContainsKey(s[i]) ){
                    int left = 0;
                    int pos = 0;
                    for(int j = i; j < s.Length; j++){
                        if( s[j] == s[i] ) left++;
                        else if(s[j] == map[s[i]]) {
                            left--;
                            pos = j;
                        }
                        if(left == 0){
                            break;
                        }
                    }
                    if(left == 0){
                        if(i+1<pos) expresion = s.Substring(i+1, pos-(i+1));
                        valido = validar(expresion, ref map);
                        i = pos;
                    }
                    if(!valido) return false;
                }
            }
            return valido;
        }
    }
}
