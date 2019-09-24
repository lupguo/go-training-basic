package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	reverseList(Head)
}

var Head *ListNode

func reverseListTraversal(node *ListNode) *ListNode {
	if node != nil {
		cnode := reverseListTraversal(node.Next)
		if cnode == nil {
			Head = node
		} else {
			cnode.Next = node
			node.Next = nil
		}
	}
	return node
}

func reverseList(head *ListNode) *ListNode {
	Head = nil
	reverseListTraversal(head)
	return Head
}
