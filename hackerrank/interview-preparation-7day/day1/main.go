package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	testCase_findMedian()
}

//day 1

func findTheMedian(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	n := len(arr)
	mid := n / 2
	// 1 2 3
	// 1 2 3 4
	if n%2 == 0 {
		return (arr[mid-1] + arr[mid]) / 2
	}
	return arr[mid]
}

//mock testCase
func testCase_findMedian() {
	arr := []int32{5, 3, 1, 2, 4}
	fmt.Println(findTheMedian(arr)) // 3

	arr = []int32{1, 2, 3, 4}
	fmt.Println(findTheMedian(arr)) // 2
}

func timeConv(s string) string {
	if s[:2] == "12" && s[8:] == "PM" {
		return s[:8]
	}
	if s[:2] == "12" && s[8:] == "AM" {
		return "00" + s[2:8]
	}
	if s[8:] == "PM" {
		hour, _ := strconv.Atoi(s[:2])
		hour += 12
		return strconv.Itoa(hour) + s[2:8]
	}
	return s[:8]
}

func testCase_timeConv() {
	s := "11:01:00PM" // : initialize value
	println(timeConv(s))

	s = "11:01:00AM" // = reassign value
	println(timeConv(s))
}

func miniMaxSum(arr []int32) {
	var min, max int64 // type val must be same

	// int 32 cant't hold value
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	for i, v := range arr { //  _ can't use as val
		if i > 0 {
			max += int64(v)
		}
		if i < 4 {
			min += int64(v)
		}
	}
	fmt.Println(min, max)
}

func testCase_miniMaxSum() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr) // expected 2063136757 2744467344
}

func plusMinus(arr []int32) {
	var pos, neg, zero float32
	var n = float32(len(arr))
	for _, v := range arr { // index, value
		if v > 0 {
			pos += 1
		} else if v < 0 {
			neg += 1
		} else if v == 0 {
			zero += 1
		}
	}

	fmt.Printf("%f\n%f\n%f", pos/n, neg/n, zero/n)
}

func testCase_pluMinus() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr) // 0.500000, 0.333333, 0.166667
}
