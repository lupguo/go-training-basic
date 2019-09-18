package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		// empty list
		return nil
	} else if head.Next == nil {
		// only one element
		return head
	} else {
		// loop the head list into stack, dump to newHead list
		var stackNodes []*ListNode
		for n := head; n != nil; n = n.Next {
			stackNodes = append(stackNodes, n)
		}
		// new head link list
		newHead := new(ListNode)
		n := newHead
		for i := len(stackNodes) - 1; i > -1; i-- {
			n.Next = &ListNode{
				Val:  stackNodes[i].Val,
				Next: nil,
			}
			n = n.Next
		}
		return newHead.Next
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// input
	head := &ListNode{
		Val:  -1,
		Next: new(ListNode),
	}
	n := head
	for _, v := range nums {
		n.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		n = n.Next
	}
	head = head.Next
	// reversal
	reversal := reverseList(head)
	// output
	for n := reversal; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
}
