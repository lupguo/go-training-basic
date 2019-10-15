package main

import "fmt"

// 25
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 示例 :
//
// 给定这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
// 说明 :
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 检测head链表是否少于k个元素，少于直接返回
	count := 0
	for cur := head; cur != nil && count < k; cur = cur.Next {
		count++
	}
	if count < k {
		return head
	} else {
		count = 0
	}

	// 反转前K个节点(head指针:链表拼接\cur指向当前指针:节点反转用\prev:节点反转用，同时指向反转链表头部哨兵节点\count节点计数器)
	var prev, next *ListNode
	cur, next := head, head
	for next != nil && count < k {
		next = next.Next
		cur.Next, prev, cur = prev, cur, next
		count++
	}

	// 第一组K元素的最后一个元素的下一个节点，指向下一个链表的prev头节点(递归得到)
	if next != nil {
		head.Next = reverseKGroup(next, k)
	}

	// 返回prev头节点指针
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	printList := func(n *ListNode) {
		for ; n != nil; n = n.Next {
			fmt.Printf("%d ", n.Val)
		}
	}

	printList(list)
	fmt.Println()
	printList(reverseKGroup(list, 2))
}
