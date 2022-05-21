package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//testCase_gridChallenge()
	//testCase_superDigit()
	//testCase_minimumBribes()
	//testCaseInput_minimumBribes()
	testCaseInput_truckTour()
	//minimimunBribesHackerrank([]int32{1,2,3,4,5})
}

//day 4

//mock
func testCaseInput_truckTour() {
	var t, n int32

	fmt.Scanf("%d", &t)

	arr := make([][]int32, t)
	for i := 0; i < int(t); i++ {
		arr[i] = make([]int32, 2)
		for j := 0; j < 2; j++ {
			fmt.Scanf("%d", &n)
			arr[i][j] = n
		}
	}
	fmt.Println(truckTour(arr)) //expected 573
	//fmt.Println(arr)

}

func testCase_gasStation() {
	arr := [][]int32{
		{1, 5}, {3, 4}, {4, 1}, {10, 3},
	}
	fmt.Println(truckTour(arr))
}

func truckTour(t [][]int32) int32 {
	var remainder int32 = 0
	var candidate int32 = 0
	var i int32

	for i = 0; i < int32(len(t)); i++ {
		remainder += t[i][0] - t[i][1]
		if remainder < 0 {
			candidate = i + 1
			remainder = 0
		}
	}
	return candidate
}

func testCaseInput_minimumBribes() {
	var t, n, elm int32

	fmt.Scanf("%d", &t)
	for i := 0; i < int(t); i++ {
		fmt.Scanf("%d", &n)
		arr := make([]int32, n)
		for j := 0; j < int(n); j++ {
			fmt.Scanf("%d", &elm)
			arr[j] = elm
		}
		minimimunBribesHackerrank(arr)
		//fmt.Println(arr)
	}

}

func testCase_minimumBribes() {
	//a := []int32{1,2,3,5,4,6,7,8}
	a := []int32{1, 2, 5, 3, 4, 7, 8, 6} //expected 8
	minimumBribes(a)
	a = []int32{5, 1, 2, 3, 7, 8, 6, 4}
	minimumBribes(a)

}

func minimimunBribesHackerrank(q []int32) {
	/* The key is to iterate starting from the end of the
	array and simultaneously rembering the minimal element so far.  */
	var bribes int32
	var last int32 = q[len(q)-1]

	for i := len(q) - 2; i >= 0; i-- { //O(n)
		if q[i] > int32(i+3) {
			fmt.Println("Too chaotic")
			return
		} else if q[i] == int32(i+3) {
			bribes += 2
		} else if q[i] > last {
			bribes++ // if q greater than before we up
		} else {
			last = q[i] // store last elm
		}
	}
	fmt.Println(bribes)
}

// this is didn't optimize
func minimumBribes(q []int32) {
	// Write your code here6
	// []int32{1,2,3,5,4,6,7,8}
	// fmt.Println(q)
	bribes := make(map[int32]int)
	for i := 0; i < len(q)-1; i++ { //O(n^2)
		for j := i + 1; j < len(q); j++ {
			if q[i]-q[j] > 2 {
				fmt.Println("Too chaotic")
				return
			}
			if q[i] > q[j] {
				q[i], q[j] = q[j], q[i]
				bribes[q[i]]++
			}
		}
	}
	var res int32
	for _, v := range bribes {
		if v > 2 {
			fmt.Println("Too chaotic")
			return
		}
		res += int32(v)
	}

	// fmt.Println(res,bribes,q)
	fmt.Println(res)

}

func testCase_superDigit() {
	/* fmt.Println(superDigit("9875",4)) */
	n := "861568688536788"
	res := superDigit(n, 100000)
	fmt.Println(res)
	res = superDigit(n, 2)
	fmt.Println(res)

	/* if we concat string with 10^5 error will occur */
	//test("861568688536788", (100000%9)+1)
}

func superDigit(n string, k int32) int32 {

	//first we count sum digit of n string
	//find modulo with nine because we only return one digit
	if len(n) == 1 {
		res, _ := strconv.Atoi(n)
		res *= int(k)
		if res%9 == 0 {
			return 9
		}
		return int32(res % 9)
	}

	var res int
	for _, v := range n {
		num, _ := strconv.Atoi(string(v))
		res += num
	}

	return superDigit(strconv.Itoa(res), k)

}

func testCase_gridChallenge() {
	arr1 := []string{
		"abc", "def", "rpl", "jkg",
	}

	fmt.Println(gridChallenge(arr1))

}

func gridChallenge(grid []string) string {
	for i := 0; i < len(grid); i++ { //nlogn
		grid[i] = sortString(grid[i])
	}

	/*
	* 1 2 3 4  "abc","def","rpl","jkg"
	* 5 6 7 4   def
	* 3 4 5 6   rpl
	*           jkg
	* */

	for row := 0; row < len(grid[0]); row++ {
		for column := 0; column < len(grid)-1; column++ {
			if grid[column][row] > grid[column+1][row] {
				return "NO"
			}
		}
	}
	return "YES"

}

//utils
func sortString(arr string) string {
	arr1 := []rune(arr)
	sort.Sort(sortByString(arr1))

	return string(arr1)
}

type sortByString []rune

func (a sortByString) Len() int           { return len(a) }
func (a sortByString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByString) Less(i, j int) bool { return a[i] < a[j] }
