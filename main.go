package main

import (
	"GoBaseToolkit/internal/collections"
	"GoBaseToolkit/internal/stats"
	"GoBaseToolkit/internal/texutil"
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 4, 5, 3, 6, 7}
	fmt.Println(collections.IndexOf(test, 9), collections.IndexOf(test, 3))
	fmt.Println(collections.RemoveDuplicates(test))
	fmt.Println(collections.Reverse(test))
	fmt.Println(collections.Filter(test, func(x int) bool { return x%2 == 0 }))
	fmt.Println(collections.Chunk(test, 3))

	fmt.Println(stats.Freq(test))
	fmt.Println(stats.UniqueCounter(test))
	fmt.Println(stats.Top(test))
	fmt.Println(stats.GroupBy(test, func(x int) string {
		if x%2 == 0 {
			return "even"
		} else {
			return "odd"
		}
	}))
	fmt.Println(stats.CountBy(test, func(x int) string {
		if x%2 == 0 {
			return "even"
		} else {
			return "odd"
		}
	}))

	fmt.Println(texutil.RunesCounter("hello世"))
	fmt.Println(texutil.ReveseString("hello世"))
	fmt.Println(texutil.IsPalindrome("топот"))
	fmt.Println(texutil.ToCamelCase("привет миру!"))
	fmt.Println(texutil.Truncate("привет миру!", 15))
	fmt.Println(texutil.Truncate("привет миру!", 5))
}
