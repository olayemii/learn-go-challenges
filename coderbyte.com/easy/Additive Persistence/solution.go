package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	Using the Go language, have the function AdditivePersistence(num) take the num parameter being passed
	which will always be a positive integer and return its additive persistence
	which is the number of times you must add the digits in num until you reach a single digit.
	For example: if num is 2718 then your program should return 2 because 2 + 7 + 1 + 8 = 18 and 1 + 8 = 9 and you stop at 9.
*/

func AdditivePersistence(num int) (times int) {
	times = 0
	for num > 10 {
		times++
		num = SumArray(GetArrayNum(num))
	}

	return times

}

func SumArray(arr []int) (sum int) {
	sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func GetArrayNum(num int) (ret []int) {
	strNum := strconv.Itoa(num)

	for _, val := range strings.Split(strNum, "") {
		j, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		ret = append(ret, j)
	}

	return ret
}

func main() {
	fmt.Print(AdditivePersistence(2718))
}
