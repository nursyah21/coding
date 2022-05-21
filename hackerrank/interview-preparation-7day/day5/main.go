package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	//fmt.Println("test")
	//testCase_queueUsingTwoStack()
	//testCase_BalancedBrackets()
	//testCaseInput_BalancedBrackets()
	//testCase_pairs()
	testCaseInput_pairs()
}

//day 5
//mock test
func testCaseInput_pairs() {
	var lenArr, target, elm int32
	fmt.Scanf("%d %d", &lenArr, &target)

	arr := make([]int32, lenArr)

	for i := 0; i < int(lenArr); i++ {
		fmt.Scanf("%d", &elm)
		arr[i] = elm
	}
	println(pairs(target, arr))
}

func testCase_pairs() {
	arr := []int32{1, 5, 3, 4, 2}
	var k int32 = 2

	println(pairs(k, arr))
}

// no optimation
func pairs(k int32, arr []int32) int32 {
	// Write your code here
	var res int32
	// 1-5,5-3,3-4,4-2

	//1-2,5-2,3-2,4-2,2-2
	//1,5,3,4,2
	//-1,3,1,2,0
	//-1,0,1,2,3
	//1,2,3,4,5

	diff := make([]int32, len(arr))

	for i, v := range arr {
		diff[i] = v - k
	}

	// optimation with sort
	sort.Slice(diff, func(a, b int) bool { return diff[a] < diff[b] })
	sort.Slice(arr, func(a, b int) bool { return arr[a] < arr[b] })
	fmt.Println(diff)
	fmt.Println(arr)

	//check
	check := func(k int32, arr []int32) bool {
		for _, v := range arr {
			if v == k {
				return true
			}
		}
		return false
	}

	for _, v := range diff {
		if check(v, arr) {
			res++
		}
	}
	return res
}

func testCaseInput_BalancedBrackets() {
	var t int
	var test string
	fmt.Scanf("%d", &t)
	count := 1
	for ; t != 0; t-- {
		fmt.Scanf("%s", &test)
		fmt.Printf("Case-%d: ", count)
		count++
		fmt.Println(balancedBrackets2(test))
	}
}

func testCase_BalancedBrackets() {
	//fmt.Println(isBalanced2("{{}}"))
	fmt.Println(balancedBrackets2("([{([]({}{}))[]}[{}]])()(()){}[()[]()]()()()[]{}{}()()()[()][[]([])()()[]]([]){}[{[][]}{}]{}{}[][]{}[](())(())()[({{}{[(())()[]]{[{}{()}{}]{[[]][[]]}()[]}}()(){[[]{}][()]()[{({})}][[([]()[{[]({(())()})}][]({})[])([{}]{()})]]}}){}{{()}}{[][[{}]][()[()]({[]{}(())})]}][{}[{}]]()()(())[[]]{[[{()}]]({})[]}({}{(([]))}[]){}[()(()[])]{}(({}))[][{}](){()}[]{({}([][][[]]))}[][((()[]({{}}[[{{[(({()}({}[[]][{([()])}({})][]{[[]{}]})())){}]{}}{()}{[][]}{}{}}][]([{}[[]{}({([]()(()()[]))})[]()]][()](()))[]]{}[[{[][]}[[[]]()[(){}][{[]{}}[{{{[]}}{}}][]([[{{()[]}[{[][{[[{[{}[]()[]]}{{}}{}]]}]}[]]}{}]][[]][{({}[])[[[{}][]]()[]]}{{[{}][]({}([][]{()}[()]){{}{}})}{}{}(([[]]()[]))()}][][{[({})[[]([[{[]}()]](([[]{}]{})))](){}[{}][]{[]}{[]([({{{}()}{[]}((){}{})}[[][]]{}[])]{})}]}{([()()[]][])}()([])])][([[]]()[])([[][]]){}[{[((([]){(){({[]})()}})){{}[()({({})}[[{[]{}}({{}[{{}{}}()]}){}]])]()}{}]{{}}}[]()]]{}{}]]]){}){}){()}()[]({}())]{()[]}[]{}{[]}(){[][[][]{}]}[{}{[{}{}]}]{[{}({}[()({{}})])()()]({[]}()((())))}"))

	/* printlist(l1)
	print(l1.Front().Value.(string))
	print(l1.Back().Value.(string))  */
	print(balancedBrackets2("()"))
}

func balancedBrackets2(s string) string {
	if len(s)%2 != 0 {
		return "NO"
	}
	//print(len(s))

	stack := list.New()
	pair := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	counter := 0
	for _, v := range s {
		counter++
		v_ := string(v)
		if v_ == "{" || v_ == "[" || v_ == "(" {
			stack.PushFront(v_)
		} else if stack.Len() != 0 {
			delElm := stack.Front()
			if delElm.Value.(string) == pair[v_] {
				stack.Remove(delElm)
			} else {
				print(counter)
				return "NO"
			}
		} else {
			return "NO"
		}
	}

	//print(counter)
	if stack.Len() == 0 {
		return "YES"
	}

	return "NO"
}

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}
func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func BalancedBrackets(s string) string {
	if len(s)%2 != 0 {
		return "NO"
	}
	stack := list.New()
	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			stack.PushFront(val)
		} else if val == ')' {
			if stack.Front().Value.(rune) != '(' {
				return "NO"
			}
			delElm := stack.Front()
			stack.Remove(delElm)
		} else if val == ']' {
			if stack.Front().Value.(rune) != '[' {
				return "NO"
			}
			delElm := stack.Front()
			stack.Remove(delElm)
		} else if val == '}' {
			if stack.Front().Value.(rune) != '{' {
				return "NO"
			}
			delElm := stack.Front()
			stack.Remove(delElm)
		}
	}
	if stack.Len() == 0 {
		return "YES"
	}
	return "NO"
}

func testCase_queueUsingTwoStack() {
	queueUsingTwoStack()
}

func queueUsingTwoStack() {
	//1 x: enqueue elm x
	//2: Dequeue front elm
	//3: print front elm
	var t, case_t, elm int
	fmt.Scanf("%d", &t)

	l1 := list.New()

	for ; t != 0; t-- {
		fmt.Scanf("%d", &case_t)
		switch case_t {
		case 1:
			fmt.Scanf("%d", &elm)
			l1.PushBack(elm)
		case 2:
			del := l1.Front()
			l1.Remove(del)
		case 3:
			fmt.Println(l1.Front().Value.(int))
		}

	}
	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ")
	}
}

func testCase_mergeSortedLink() {
	list1 := list.New()
	list1.PushBack(5)
	list1.PushBack(1)
	list1 = sortLinkedList(list1)
	println("list1:")
	printlist(list1)

	list2 := list.New()
	list2.PushBack(4)
	list2.PushBack(3)
	list2 = sortLinkedList(list2)
	println("list2 :")
	printlist(list2)

	res := mergeTowSortedLink(list1, list2)

	println("merge list: ")
	printlist(res)
}

func printlist(head *list.List) {
	for i := head.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/* bubble sort */
func sortLinkedList(node *list.List) *list.List {
	current := node.Front()
	if node.Front() == nil {
		return nil
	} else {
		for ; current != nil; current = current.Next() {
			index := current.Next()
			for ; index != nil; index = index.Next() {
				/* swap current and next current */
				if current.Value.(int) > index.Value.(int) {
					current.Value, index.Value = index.Value, current.Value
				}
			}
		}
	}
	return node
}

func mergeTowSortedLink(headA, headB *list.List) *list.List {
	node1 := headA.Front()
	node2 := headB.Front()
	resNode := list.New()

	for node1 != nil && node2 != nil {
		if node1.Value.(int) < node2.Value.(int) {
			resNode.PushBack(node1.Value)
			node1 = node1.Next()
		} else {
			resNode.PushBack(node2.Value)
			node2 = node2.Next()
		}
	}
	for node1 != nil {
		resNode.PushBack(node1.Value)
		node1 = node1.Next()
	}
	for node2 != nil {
		resNode.PushBack(node2.Value)
		node2 = node2.Next()
	}

	return resNode

}
