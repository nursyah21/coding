package main

import (
	"C"
	"fmt"
	"sort"
	"strings"
	"time"
)

//utils
func swap[T int | int32](a *T, b *T) {
	*a, *b = *b, *a
}

func delChar(st string, idx int) string {
	s := []rune(st)
	return string(append(s[0:idx], s[idx+1:]...))
}

func main() {
	//testCase_towerBreakers()
	//testCase_palindromeIndex()
}

// day 3

//mock test

func testWithTime(f func(s string, v int32), s string, v int32) {
	st := time.Now().Nanosecond()
	f(s, v)
	ed := time.Now().Nanosecond()
	fmt.Println(ed - st)
}

var t int32 = 1

func run_palindromeIndex(s string, v int32) {

	if palindromIndexHackerrank(s) == v {
		fmt.Printf("testCase%d: success\n", t)
	} else {
		fmt.Printf("testCase%d: failed\n", t)
	}
	t += 1
}

func run_palindromeIndexaNoHackkerank(s string, v int32) {

	if palindromeIndex(s) == v {
		fmt.Printf("testCase%d: success\n", t)
	} else {
		fmt.Printf("testCase%d: failed\n", t)
	}
	t += 1
}

func testCase_palindromeIndex() {
	st := time.Now().Local().Nanosecond()
	run_palindromeIndexaNoHackkerank("aabfpefmepompwivwpvimewpvervpmepbmeprmpwbmepbmeimbiermbprovmrev,[swpvmepbmewbie", 2)
	ed := time.Now().Local().Nanosecond()
	fmt.Printf("%dns %fms\n\n", (ed - st), float64((ed-st))/100000)

	st = time.Now().Local().Nanosecond()
	run_palindromeIndexaNoHackkerank("aabfpefmepompwivwpvimewpvervpmepbmeprmpwbmepbmeimbiermbprovmrev,[swpvmepbmewbie", 2)
	ed = time.Now().Local().Nanosecond()
	fmt.Printf("%dns %fms\n\n", (ed - st), float64((ed-st))/100000)

	//fmt.Println()
	st = time.Now().Nanosecond()
	run_palindromeIndex("aabfpefmepompwivwpvimewpvervpmepbmeprmpwbmepbmeimbiermbprovmrev,[swpvmepbmewbie", 2)
	ed = time.Now().Nanosecond()
	fmt.Printf("%dns %fms\n\n", (ed - st), float64((ed-st))/1000000)

	st = time.Now().Nanosecond()
	run_palindromeIndexaNoHackkerank("aabfpefmepompwivwpvimewpvervpmepbmeprmpwbmepbmeimbiermbprovmrev,[swpvmepbmewbie", 2)
	ed = time.Now().Nanosecond()
	fmt.Printf("%dns %fms\n\n", (ed - st), float64((ed-st))/1000000)

}

func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		} //logn
	}

	return true
}

//this solution from hackerrank
func palindromIndexHackerrank(s string) int32 {
	i, j := 0, len(s)-1

	if checkPalindrome(s) {
		return -1
	} // 1
	// tamat
	for (i < j) && (s[i] == s[j]) {
		i += 1
		j -= 1
	}

	if i == j {
		return 0
	}

	if checkPalindrome(s[i+1 : j+1]) { //mat
		return int32(i)
	}
	if checkPalindrome(s[i:j]) { //ama
		return int32(j)
	}
	return -1
}

// my solution didn't optimize
func palindromeIndex(s string) int32 {
	// Write your code here
	//lowercase
	//make index character to remove to make palindrome

	// if you make delChar func in here, the return only change once

	var res int32 = -1

	if checkPalindrome(s) {
		return res
	} // if already palindrome return -1
	for i := range s { //n
		if checkPalindrome(delChar(s, i)) { // false //logn
			res = int32(i)
		}

	}

	return res
}

func testCase_caesarChiper() {
	fmt.Println(caesarCipher("Zero-zero", 26))
	fmt.Println(caesarCipher("middle-Outz", 83))
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	var res []string
	//fmt.Printf("%d %d %d %d %d",rune('-'), rune('A'),rune('Z'), rune('a'),rune('z')) //know rune - A Z
	for _, v := range []rune(s) {

		if v >= 65 && v <= 90 { // A-Z
			v += k % 26
			if v > 90 {
				v -= 26
			} // back to beginning
		}
		if v >= 97 && v <= 122 { // a-z
			v += k % 26
			if v > 122 {
				v -= 26
			} // back to beginning
		}

		res = append(res, string(v))
	}
	return strings.Join(res, "")
}

//skip towerBreakouts
func testCaseStdOut_towerBreakers() {
	var t, m, n int32
	fmt.Scanf("%d", &t)

	for ; t != 0; t-- {
		fmt.Scanf("%d %d", &m, &n)
		fmt.Println(towerBreakers(m, n))
	}
}

func testCase_towerBreakers() {
	fmt.Println(towerBreakers(2, 2))      //expected 2
	fmt.Println(towerBreakers(1, 4))      //expected 1
	fmt.Println(towerBreakers(100000, 1)) //expected 2
	fmt.Println(towerBreakers(100001, 1)) //expected 2
}

//this code not optimized, thought you get correct
// will failed if we run 374625 796723
func towerBreakers(n int32, m int32) int32 { // n: num of tower, m: height of each tower
	// imagine 2 players p1, and p2 to destroy tower
	// the code absolutely easy
	// but to understanding this problem quite difficult
	var win int32 = 1
	if n%2 == 0 {
		win = 2
	} else {
		win = 1
		if m == 1 {
			win = 2
		}
	}

	//rules of problem
	/*
		- Initially there are `n` towers.
		- Each tower is of height `m`

		- The players move in alternating turns.
		- In each turn, a player can choose a tower of height `x`
		  and reduce its height to `y`, where `1 <= y < x` and `y` evenly divides `x`
		.
		- If the current player is unable to make a move, they lose the game.
	*/
	/*
			Even though I got the answer and the underlying logic, I still think this is actually a very tough problem.
			I say so because of the following reasons:

		    You have to first understand the complicated game and its rules properly (slightly challenging)

		    You then have to find the correct way to win this game from the perspective of each player (extremely challenging)

			So basically, there are two cases here. Both cases rely heavily on the idea that an even number of towers (say 2n)
			can be thought of as a collection of n pairs of towers.

			CASE 1) If there are an even number of towers Player 1 will go first, and Player 2 will basically copy player
			1. Whatever player 1 does to a tower, player 2 will do the exact same thing to the other tower belonging to the same pair.
			In this case Player 2 will always win. // BECAUSE THEY MUST HAVE SAME EFFECTIVE ThOUGH TO DESTROY TOWER
			// AND PLAYER1 IS THE LAST PLAYER

			CASE 2) If there are an odd number of towers Player 1 will go first, and simply destroy any single tower by setting it equal to 1.
			Now we again have the same situation from CASE 1 but this time the roles of both players have been reversed.
			In this case Player 1 will always win.

			Finally, there is also the trivial case where n does not matter because m = 1. In this case Player 2 will obviously always win.
	*/
	return win
}

func testCase_zigzagSequence() {
	arr1 := []int32{1, 2, 3, 4, 5}
	zigzagSequence(arr1) //expected 1 2 5 4 3
}

func zigzagSequence(arr []int32) {
	sort.Slice(arr, func(a1, a2 int) bool { return arr[a1] < arr[a2] }) //sort array
	n := len(arr)
	mid := (n+1)/2 - 1
	swap(&arr[mid], &arr[n-1]) //swap mid arr with last arr

	st := mid + 1
	ed := n - 2
	for st <= ed { //iterate from mid to last
		swap(&arr[st], &arr[ed])
		st += 1
		ed -= 1
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
	}
}
