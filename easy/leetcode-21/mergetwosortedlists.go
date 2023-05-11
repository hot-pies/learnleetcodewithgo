package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var fl1 = &ListNode{Val: 1}
	var fl2 = &ListNode{Val: 2}
	var fl3 = &ListNode{Val: 4}

	var sl1 = &ListNode{Val: 1}
	var sl2 = &ListNode{Val: 3}
	var sl3 = &ListNode{Val: 4}

	fl1.Next = fl2
	fl2.Next = fl3

	sl1.Next = sl2
	sl2.Next = sl3

	fmt.Println(mergeTwoLists(fl1, sl1))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		var dummy = &ListNode{}
		return dummy
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list2 != nil && list1 == nil {
		return list2
	}

	mergedlist := &ListNode{}

	current := mergedlist

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			current.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	// for mergedlist.Next != nil {
	// 	fmt.Println("Print merge list ", mergedlist.Val)
	// 	mergedlist = mergedlist.Next
	// }

	return mergedlist.Next

}
