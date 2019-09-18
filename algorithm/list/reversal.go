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
	} else if (head.Next == nil) {
		// only one element
		return head
	} else {
		// loop the head list into stack, dump to newHead list
		var stackNodes []*ListNode
		for head.Next != nil {
			stackNodes = append(stackNodes, head)
		}
		newHead := new(ListNode)
		for _, node := range stackNodes {
			newHead.Next = node
		}
		return newHead
	}
}

func main() {
	vals := []int{1, 2, 3, 4, 5}

	// input
	head := &ListNode{
		Val:  0,
		Next: new(ListNode),
	}
	for _, v := range vals {
		head.Next = &ListNode{
			Val:  v,
			Next: new(ListNode),
		}
	}

	// reversal
	reversal := reverseList(head)

	// output
	for n := reversal; n != nil && n.Next != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
}
