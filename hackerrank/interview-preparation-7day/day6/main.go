//go:build windows
// +build windows

package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"github.com/interview_preparation/Main/graphs"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	graphs.TestGraph()

	//testCaseInput_simpleTextEditor()
	//testCaseInput_legoBlocks()
	//fmt.Println(coinChange(12))
	//testCaseInput_cookies()
	//test()
	//cookiesHeap(1, []int32{1, 3,2})
	/* nums := []int{3,2,10,20,4,60,12}
	h := &IntHeap{}

	for _,v := range nums{
		heap.Push(h, v)
	}
	for i := 0; i < len(nums); i++{
		fmt.Printf("%d, ", heap.Pop(h).(int))
	}
	fmt.Println("") */

}
func testCaseInput_cookies() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	//defer writer.Flush()

	var n, k int32
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	A := make([]int32, n)

	var elm int32

	for i := 0; i < int(n); i++ {
		fmt.Fscanf(reader, "%d ", &elm)
		A[i] = elm
	}
	//println(A[0])
	//fmt.Fprintf(writer,"%d\n", cookiesHeap2(k, A))
	fmt.Println(cookiesHeap2(k, A))
	//fmt.Println("non heap:")
	//fmt.Println(cookies(k,A))
}

func cookiesHeap2(k int32, A []int32) int32 {
	//O nlog(n)
	sort.Slice(A, func(a, b int) bool { return A[a] < A[b] }) // must sorting before insert to heap

	h := &IntHeap{}
	for _, v := range A {
		heap.Push(h, v)
	}

	fmt.Println(h)
	var elm1, elm2, newElm int32
	for i := 1; h.Len() > 0; i++ {
		heap.Fix(h, h.Len()-1)
		elm1 = h.Pop().(int32)
		heap.Fix(h, h.Len()-1)

		newElm += elm1
		if h.Len() > 0 {
			elm2 = h.Pop().(int32)
			heap.Fix(h, h.Len()-1)
		}
		newElm += 2 * elm2
		heap.Push(h, newElm)
		heap.Fix(h, h.Len()-1)

		fmt.Println(i, elm1, elm2, newElm)
		if h.Peek().(int32) >= k {
			break
		}
	}
	fmt.Println(h)

	var lowest int32 = 0

	/* for h.Peek().(int32) < k{ // don't use h.Len(), because if we pop, it will reduce
		fmt.Print(h)
		//heap.Fix(h, h.Len()-2)
		elm1 := h.Pop().(int32)
		elm2 := h.Pop().(int32)
		newElm := elm1 + 2* elm2
		fmt.Printf("\n %d %d %d\n", elm1, elm2, newElm)
		heap.Push(h, newElm)
		if h.Len() < 2{
			lowest = -1
			break
		}
		lowest++
	}
	fmt.Println(h, k) */
	return lowest
}

func test() {
	insert := func(k int32, arr []int32) {
		/* old := make([]int32, len(arr[2:]))
		for i:=0;i < len(arr[2:]);i++{
			old[i] = arr[2+i]
		} */
		old := []int32{4, 5}
		fmt.Println(old)
		arr = append(arr[:2], k)
		arr = append(arr[0:], old...)
		fmt.Println(old)

	}
	A := []int32{1, 2, 4, 5}
	insert(3, A)

	fmt.Println(A)
}

//2 7 3 6 4 6
//2 3 4 6 6 7
//8 4 6 6 7
//8 16 6 7
//8 16
// need improve
func cookies(k int32, A []int32) int32 {

	// Write your code here
	swapA := func(A []int32) {
		for i := 0; i < len(A)-1; i++ {
			if A[i] > A[i+1] {
				A[i], A[i+1] = A[i+1], A[i]
			} else {
				break
			}
		}
	}
	/* var mysum int32 = 0
	for _,v := range A{
		mysum += v
	} */

	sort.Slice(A, func(a, b int) bool { return A[a] < A[b] })
	var res int32 = 0
	n := len(A)
	for k > A[0] {
		//println(res, A[0], k, mysum)
		newA := A[0] + 2*A[1]
		fmt.Println(A, "\n", A[0], A[1], newA)
		A = append(A[1:]) //remvove 2 elm ,,A[0:] remove first elm
		A[0] = newA
		n--
		if n < 2 {
			res = -1
			break
		}
		res++
		swapA(A)
		//sort.Slice(A, func(a,b int) bool {return A[a] < A[b]})
	}
	fmt.Println(A, k)

	return res
}

type IntHeap []int32

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int32))
}

//see first elm Heap
func (h *IntHeap) Peek() interface{} {
	//n := len(*h)
	old := *h
	x := old[0]
	return x
}
func (h *IntHeap) Pop() interface{} { // pop first elm
	old := *h
	//n := len(old)
	x := old[0]
	*h = old[1:]
	return x
}

type MaxHeap struct {
	array []int32
}

func (h *MaxHeap) Insert(key int32) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int32 {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		//cannot extract if h.array empty
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parrent(index)] > h.array[index] {
		h.swap(parrent(index), index)
		index = parrent(index)
	}
}

func parrent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func cookiesHeap(k int32, A []int32) int32 {
	//h := heap.Init()
	m := &MaxHeap{}
	test := []int32{7, 1, 3, 5, 4, 10}
	for _, v := range test {
		m.Insert(v)
	}

	fmt.Println(m)
	fmt.Println(m.Extract())
	fmt.Println(m.Extract())
	fmt.Println(m.Extract())
	fmt.Println(m)
	return 0
}

func testCaseInput_legoBlocks() {
	defer writer.Flush()

	var n, h, w, res int32
	fmt.Fscanf(reader, "%d\n", &n)

	for ; n != 0; n-- {
		fmt.Fscanf(reader, "%d %d\n", &h, &w)
		res = legoBlocks(h, w)
		fmt.Fprintf(writer, "%d\n", res)
	}
}

// need solve another problem
// this is dp prblem
// syill didn't understand the problem
func legoBlocks(n int32, m int32) int32 {
	// Write your code here
	// n height m width
	var upper float64 = 1000000007 //prevent overflow
	a := []float64{0, 1, 2, 4, 8}  //wall layouts with width i

	for j := 5; j <= int(m); j++ {
		res := math.Mod((a[j-1] + a[j-2] + a[j-3] + a[j-4]), upper)

		a = append(a, res)
	}
	for i := 0; i <= int(m); i++ {
		res := math.Pow(a[i], float64(n))
		res = math.Mod(res, upper)
		a[i] = res // number possible layouts
	}

	println(a, a[0])

	return int32(upper)
}

func coinChange(n int32) int32 {
	coin := []int32{1, 5, 10}
	ways := make([]int32, n)

	ways[0] = 1
	for i := 0; i < len(coin); i++ {
		for j := 0; j < int(n); j++ {
			if int32(j) >= coin[i] {
				ways[j] += ways[int32(j)-coin[i]]
			}
		}
	}

	return ways[int(n)-1]

}

func testCaseInput_simpleTextEditor() {
	simpleTextEditor2()
}

//replace scanf with new Reader
// speed improve from

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

// very fast, we use bufio to make faster 8x
func simpleTextEditor2() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	/* 	1. append (W)- Append string W to the end of S
	.		2. delete (K)- Delete the last K characters of S.
			3. print (K)- Print the Kth character of S
			4. undo - Undo the last (not previously undone)
			operation of type 1 or 2,  */

	defer writer.Flush()

	var n int32
	// improve
	fmt.Fscanf(reader, "%d\n", &n)

	var res, newres string

	oldres := list.New()

	var Case int
	for ; n != 0; n-- {
		fmt.Fscanf(reader, "%d %s\n", &Case, &newres)
		switch Case {
		case 1:
			oldres.PushFront(res)
			res += newres
		case 2:
			oldres.PushFront(res)
			delInt, _ := strconv.Atoi(newres)

			res = res[:len(res)-delInt]
		case 3:
			printint, _ := strconv.Atoi(newres)
			printf("%s\n", string(res[printint-1]))
		case 4:
			if oldres.Len() > 0 {
				last := oldres.Front()
				res = last.Value.(string)
				oldres.Remove(last)
			}
		}
	}
}

//this slow because scanf is parsing
//take 7.09s for 1000000 input
func simpleTextEditor() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	/* 	1. append (W)- Append string W to the end of S
	.		2. delete (K)- Delete the last K characters of S.
			3. print (K)- Print the Kth character of S
			4. undo - Undo the last (not previously undone)
			operation of type 1 or 2,  */

	var n int32
	fmt.Scanf("%d", &n)

	var res string

	oldres := list.New()

	//var overflow  int32 = 10000
	for ; n != 0; n-- {

		var Case int
		fmt.Scanf("%d", &Case)
		switch Case {
		case 1:
			oldres.PushFront(res)
			var newres string
			fmt.Scanf("%s", &newres)
			res += newres
		case 2:
			oldres.PushFront(res)
			var delInt int
			fmt.Scanf("%d", &delInt)
			res = res[:len(res)-delInt]
		case 3:
			var printint int
			fmt.Scanf("%d", &printint)
			fmt.Println(string(res[printint-1]))
		case 4:
			if oldres.Len() > 0 {
				last := oldres.Front()
				res = last.Value.(string)
				oldres.Remove(last)
			}
		}
	}
}
