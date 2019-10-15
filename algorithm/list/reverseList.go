package main

import "fmt"

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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

// 迭代反转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for p := head; head != nil; {
		// 注意链表的表头很重要，不要随意丢了
		head = head.Next
		p.Next, prev, p = prev, p, head
	}
	return prev
}

// 递归反转
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newList := reverseListRecursion(head.Next)
	// 低效的迭代，为了找到尾节点
	//for p := newList; ; p = p.Next {
	//	if p.Next == nil {
	//		p.Next = head
	//		head.Next = nil
	//		break
	//	}
	//}
	// 高效的指针引用，添加尾节点
	head.Next.Next = head
	//注意为空判断，否则会出现循环链表
	head.Next = nil
	return newList
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	fplinked := func(p *ListNode) {
		for p != nil {
			fmt.Printf("%d ", p.Val)
			p = p.Next
		}
	}

	fplinked(head)
	//fmt.Println()
	//fplinked(reverseList(head))
	fmt.Println()
	fplinked(reverseListRecursion(head))
}
