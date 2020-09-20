package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(nums []int) *ListNode {
	var head *ListNode = nil
	tail := head
	for _, v := range nums {
		p := &ListNode{Val: v}
		if head == nil {
			head = p
			tail = p
			continue
		}
		tail.Next = p
		tail = tail.Next
	}
	return head
}
func printList(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			q := p.Next
			p.Next = q.Next
		} else {
			p = p.Next
		}
	}
	return head
}
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	p := head
	for p.Next != nil && p.Next.Next != nil {
		//如果找到重复的链表段，开始删除
		if p.Next.Val == p.Next.Next.Val {
			rmVal := p.Next.Val
			q := p.Next.Next
			for q.Next != nil && q.Next.Val == rmVal {
				q = q.Next
			}
			p.Next = q.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	var curr *ListNode = head
	for curr != nil {
		p := curr.Next
		curr.Next = pre
		pre = curr
		curr = p
	}
	return pre
}
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || n-m == 0 {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	p := head
	l := 1
	for l < m {
		p = p.Next
		l++
	}
	middle := p.Next
	pre := p.Next
	curr := pre.Next
	for l < n && curr != nil {
		q := curr.Next
		curr.Next = pre
		pre = curr
		curr = q
		l++
	}
	p.Next = pre
	middle.Next = curr
	return dummy.Next
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	p := dummy
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
func mergeTwoList_recur(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoList_recur(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList_recur(l1, l2.Next)
		return l2
	}
}
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{Val: 0}
	dummy2 := &ListNode{Val: 0}
	dummy1.Next = head
	p1 := dummy1
	p2 := dummy2
	for p1.Next != nil {
		if p1.Next.Val >= x {
			p2.Next = p1.Next
			p1.Next = p1.Next.Next
			p2 = p2.Next
		} else {
			p1 = p1.Next
		}
	}
	p1.Next = dummy2.Next
	p2.Next = nil
	return dummy1.Next
}
func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	next := middle.Next
	middle.Next = nil
	left := mergeSort(head)
	right := mergeSort(next)
	return mergeList(left, right)
}
func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
}
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	middle := findMiddle(head)
	next := middle.Next
	middle.Next = nil
	//逆序后端链表
	next = reverseListRecursion(next)

	p, q := head, next
	for q != nil {
		tmp := q.Next
		q.Next = p.Next
		p.Next = q
		q = tmp
		p = p.Next.Next
	}
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow == fast {
		return true
	}
	return false
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow == fast {
		pre := head
		post := fast.Next
		for pre != post {
			pre = pre.Next
			post = post.Next
		}
		return post

	}
	return nil
}
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	middle := findMiddle(head)
	post := middle.Next
	middle.Next = nil
	pre := head
	post = reverseListRecursion(post)
	for post != nil {
		if post.Val != pre.Val {
			return false
		}
		post = post.Next
		pre = pre.Next
	}
	return true
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//存储原始节点到下表索引映射
	randomNode := make(map[*Node]int)
	//存储新节点，用于通过下标寻找random节点
	resultSlice := make([]*Node, 0)
	p := head
	result := &Node{Val: 0}
	tail := result
	idx := 0
	for p != nil {
		tmp := &Node{Val: p.Val, Random: p.Random}
		tail.Next = tmp
		randomNode[p] = idx
		resultSlice = append(resultSlice, tmp)
		idx++
		tail = tail.Next
		p = p.Next
	}
	for _, v := range resultSlice {
		if v.Random != nil {
			idx := randomNode[v.Random]
			v.Random = resultSlice[idx]
		}
	}
	return result.Next
}

func createTest() {
	fmt.Println("heelo world")
}

func main() {
	nums := []int{3, 3, 4, 34, 56, 7, 7, 8, 9, 10, 20, 20, 39, 89, 89}
	//test := []int{3,5,8,100}
	head := createList(nums)
	//head1 := createList(test)
	fmt.Println(printList(head))
	//fmt.Println(printList(deleteDuplicates(head)))
	//fmt.Println(printList(deleteDuplicatesII(head)))
	//fmt.Println(printList(reverseList(head)))
	//fmt.Println(printList(reverseListRecursion(head)))
	//fmt.Println(printList(reverseBetween(head,1,len(nums))))
	//fmt.Println(printList(reverseBetween(head1,1,2)))
	//fmt.Println(printList(mergeTwoList_recur(head,head1)))
	//fmt.Println(printList(reverse(head)))
	//fmt.Println(printList(partition(head,29)))
	//fmt.Println(printList(sortList(head)))
	reorderList(head)
	fmt.Println(printList(head))
}
