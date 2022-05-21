package main

import (
	"fmt"
	"os"

	"bufio"
	"sort"
)

func mainHeap() {
	testCaseInput_cookies()
}

func testCaseInput_cookies() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	//defer writer.Flush()

	var n, k, predict int32
	fmt.Fscanf(reader, "%d %d %d\n", &n, &k, &predict)
	A := make([]int32, n)

	var elm int32

	for i := 0; i < int(n); i++ {
		fmt.Fscanf(reader, "%d ", &elm)
		A[i] = elm
	}
	//println(A[0])
	//fmt.Fprintf(writer,"%d\n", cookiesHeap2(k, A))
	//fmt.Println("in main: ", k)
	res := cookiesHeap2(k, A)
	//resNoHeap := cookies(k,A)
	fmt.Println("heap:", res)
	if res == predict {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
	//fmt.Println(cookies(k,A))
	/* fmt.Println("non heap:", resNoHeap)
	if resNoHeap == predict {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	} */
	//fmt.Println(cookies(k,A))
}

func cookiesHeap2(k int32, A []int32) int32 {
	//O nlog(n)
	//sort.Slice(A, func(a, b int) bool {return A[a] > A[b]}) // must sorting before insert to heap

	h := &MinHeap{}

	for _, v := range A {
		h.Insert(v)
	}

	var lowest int32 = 0
	fmt.Println("before:", k)

	// i don't use heap, because get me wrong after 2xpop and push
	for h.Peek() < k { // don't use h.Len(), because if we pop, it will reduce
		//fmt.Print(h) //use h pop, and heap pop will show diff result
		elm1 := h.Extract()

		if h.Len() < 1 { //check if we can remove heap agein
			lowest = -1
			break
		}

		elm2 := h.Extract()
		newElm := elm1 + 2*elm2
		h.Insert(newElm)

		lowest++
	}
	//fmt.Println("\n", h, k)

	return lowest
}

type MinHeap struct {
	array []int32
}

func (h *MinHeap) Len() int {
	return len(h.array)
}

func (h *MinHeap) Insert(key int32) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MinHeap) Peek() int32 {
	return h.array[0]
}

func (h *MinHeap) Extract() int32 {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		return -1
	}

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxheapifyup will heapify
func (h *MinHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	var childToCompare int32 = 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = int32(l)
		} else if h.array[l] < h.array[r] {
			childToCompare = int32(l)
		} else {
			childToCompare = int32(r)
		}

		if h.array[index] > h.array[childToCompare] {
			h.swap(index, int(childToCompare))
			index = int(childToCompare)
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) swap(a, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

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
		//fmt.Println(A,"\n", A[0], A[1], newA)
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
	//sort.Slice(A, func(a,b int) bool {return A[a] < A[b]})
	fmt.Println(A, k)

	return res
}
