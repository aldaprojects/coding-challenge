/*
Merge two sorted linked lists and return it as a new list. 
The new list should be made by splicing together the nodes 
of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

*/

namespace solution {

    /**
    * Definition for singly-linked list.
    * public class ListNode {
    *     public int val;
    *     public ListNode next;
    *     public ListNode(int x) { val = x; }
    * }
    */
    public class Solution {
        public ListNode MergeTwoLists(ListNode l1, ListNode l2) {
            
            ListNode newList = new ListNode(0);
            ListNode aux = newList;
            
            
            while(l1 != null && l2 != null){
                if(l1.val < l2.val){
                    aux.next = l1;
                    l1 = l1.next;
                }
                else {
                    aux.next = l2;
                    l2 = l2.next;
                }
                aux = aux.next;
            }
            
            if(l1 != null) aux.next = l1;
            
            if(l2 != null) aux.next = l2;
            
            return newList.next;
        }
    }
}

