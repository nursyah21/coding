package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(-13597155570)

	fmt.Println(a)

	res := legoBlocks(8, 6)
	if res == 402844528 {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
	//fmt.Println(legoBlocks(4, 4))
	// why i get different result at the end
	// this is will closed, can't do it with golang
}

//mock test
// my head is hurt

//just convert code from python to golang
//even though i don't know this algorithm
func legoBlocks(h int64, w int64) int64 {
	var M float64 = 1000000007  // prevent integer overflow
	a := []int64{0, 1, 2, 4, 8} // combination wall layouts with width 1

	for i := 5; i <= int(w); i++ {
		//added new combination wall layouts w
		newres := math.Mod(float64(a[i-1]+a[i-2]+a[i-3]+a[i-4]), M)
		a = append(a, int64(newres))
	}

	for i := 0; i <= int(w); i++ {
		//update each combination wall layouts to w*h
		a[i] = int64(math.Mod(math.Pow(float64(a[i]), float64(h)), M))
	}

	goodwalls := make([]int64, w+1)
	for i := 0; i <= int(w); i++ {
		goodwalls[i] = a[i]
	}

	var res int64
	for i := 1; i <= int(w); i++ {
		for j := 1; j < i; j++ {
			a1 := math.Mod(float64(goodwalls[j]), M)
			a2 := math.Mod(float64(a[i-j]), M)
			res = int64(math.Mod(float64(a1*a2), M))

			goodwalls[i] = goodwalls[i] - res
		}
		fmt.Printf("%d %d ", int64(res), int64(goodwalls[i]))
		goodwalls[i] = int64(math.Mod(float64(goodwalls[i]), M))
	}

	fmt.Println(goodwalls)
	//big.NewInt()
	return int64(goodwalls[w])
}

// use code from editorial
// func legoBlocksHackerrank()

func pow(a int32, p int32) int32 {
	var mod int32 = 1000000007
	var ans int32
	for p > 0 {
		if p%2 != 0 {
			ans = ans * ans % mod
			p /= 2
		}
	}

	return ans
}

func coinChange(coins []int32, amount int32) int32 {
	combination := make([]int32, amount)
	combination[0] = 1

	// example combination coin: 1,2,5,10 amount: 4
	// first make array with size amount and init first array to 1, so we can do increment
	// [1,0,0,0 ]
	// combination[0] += combination[1-1] -->1
	// combination[1] += combination[2-1] -->1
	// so first iterate we have [1,1,1,1]
	// than we iterate again
	// combination[]

	for _, c := range coins { // iterate combination of coin
		for i := c; i < amount; i++ { //iterate each coin
			combination[i] += combination[i-c]
		}
	}

	fmt.Println(combination)
	return combination[amount-1]
}
