package gol

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n ListNode) ToString() string {
	p := &n
	strs := []string{}
	for {
		if p == nil {
			break
		}
		strs = append(strs, fmt.Sprintf("%d", p.Val))
		p = p.Next
	}
	return strings.Join(strs, "->")
}

func ParseList(text string) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, num := range strings.Split(strings.TrimSpace(text), "->") {
		val, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			panic(fmt.Sprintf("list node parse error: %+v", err))
		}
		p := &ListNode{
			Val:  int(val),
			Next: nil,
		}
		if head == nil {
			head = p
			tail = p
		} else {
			tail.Next = p
			tail = p
		}
	}
	return head
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head.Next
	pre := head
	for {
		if p == nil {
			break
		}
		head, pre = sortP(head, pre, p)
		p = pre.Next

	}
	return head
}

func sortP(head, pre, p *ListNode) (*ListNode, *ListNode) {
	n := head
	var nPre *ListNode
	var tail *ListNode
	for {
		if n == p {
			tail = p
			break
		}
		if n.Val > p.Val {
			if nPre == nil {
				head = p
				pre.Next = p.Next
				p.Next = n
			} else {
				nPre.Next = p
				pre.Next = p.Next
				p.Next = n
			}
			tail = pre
			break
		}
		nPre = n
		n = n.Next
	}
	return head, tail
}

func SortList(head *ListNode) *ListNode {
	// if head != nil {
	// 	fmt.Printf("SortList: %+v\n", head.ToString())
	// }

	size := listLen(head)
	if size <= 1 {
		return head
	}

	p := head
	for i := 0; i < size/2-1; i++ {
		p = p.Next
	}

	lHead, rHead := head, p.Next
	p.Next = nil

	lHead = SortList(lHead)
	rHead = SortList(rHead)

	// fmt.Printf("Merge: \n%+v\n%+v\n", lHead.ToString(), rHead.ToString())
	head = nil
	var tail *ListNode
	addNode := func(n *ListNode) {
		// if head != nil {
		// 	fmt.Printf("addNode: %+v, %+v, tail=%+v\n", head.ToString(), n.Val, tail.Val)
		// }

		if head == nil {
			head = n
			tail = n
		} else {
			tail.Next = n
			tail = n
		}
	}
	for {
		if lHead == nil {
			addNode(rHead)
			break
		}
		if rHead == nil {
			addNode(lHead)
			break
		}
		if lHead.Val < rHead.Val {
			addNode(lHead)
			lHead = lHead.Next
		} else {
			addNode(rHead)
			rHead = rHead.Next
		}
	}

	// fmt.Printf("Merged: %+v\n", head.ToString())
	return head
}

func listLen(head *ListNode) int {
	size := 0
	for p := head; p != nil; p = p.Next {
		size++
	}
	return size
}
