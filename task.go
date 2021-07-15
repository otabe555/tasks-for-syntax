package main

import (
	"fmt"
)

// I. Fibonacci
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// II. FizzBuzz
func fizzBuzz(n int) []string {
	var ans []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
		} else if i%5 != 0 && i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 && i%3 != 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, fmt.Sprint(i))
		}
	}
	return ans
}

// III. Palindrome
func isPalindrome(str string, s, e int) bool {
	if s > e {
		return true
	}
	return str[s] == str[e] && isPalindrome(str, s+1, e-1)
}
func Palindrome(str string) bool {
	return isPalindrome(str, 0, len(str))
}

// IV. Odd even sum
func oddEvenSum(arr []int) (int, int) {
	odd, even := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			odd += arr[i]
		} else {
			even += arr[i]
		}
	}
	return odd, even
}

// Has duplicate element
func hasDuplicate(arr []int) bool {
	seen := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if seen[arr[i]] {
			return true
		}
		seen[arr[i]] = true
	}
	return false
}

// the input part
func main() {

}
