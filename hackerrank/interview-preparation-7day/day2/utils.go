package main

type number interface {
	int | int32 | float64
}

// generic function
func Max[T number](a1, a2 T) T {
	if a1 > a2 {
		return a1
	}
	return a2
}

func Swap(a *int32, b *int32) {
	*a, *b = *b, *a
}

func MaxArr(arr []int32) int32 {
	res := arr[0]

	for _, v := range arr[1:] {
		if v > res {
			res = v
		}
	}

	return res
}

func MinArr(arr []int32) int32 {
	res := arr[0]

	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}

	return res
}

// must use go mod init <>
// and go run .
