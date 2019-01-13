package main

// TODO:
// 1: add support for multiple arguments
// 2: add support for a single argument to show maximum number pair

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {

	args := os.Args[1:]
	fs := [][]int{}

	for i := 0; i < len(args); i++ {
		n, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("All values must be integers")
		} else {
			fmt.Println("Number: ", n)
			f := []int{n}
			fact(&f)
			fmt.Println(f)
			fs = append(fs, f)
		}
		fmt.Println(fs)
	}
	// SCORE! Now we have a slice of slices that we can scan for commonFactors


// This is the idiot code. I hve to parse the sets by two pairs at a time. Parse
// the first two, get the result, parse the next two. Order won't matter.

	// for i := 0; i < len(fs); i++ {
	// 	fs[0] = commonFactors(fs[i], fs[i+1])
	// }
	// fmt.Println(fs)

}

func fact(f *[]int) {
	length := len(*f)
	last := (*f)[length-1]

	if length == 1 {
		*f = append(*f, 1)
		sort.Ints(*f)

		fact(f)

	} else {
		fmt.Println("Current slice: ", (*f))

		if isPrime(last) {
			// Done!
		} else {

			last := (*f)[length-1] // Last elemnt, not prime
			lpf := (*f)[length-2]  // last prime element
			*f = (*f)[:length-1]   // truncate last element

			lpf, last = getNextTwo(lpf, last) // divide last element by last prime factor

			*f = append((*f), lpf, last) // add next prime factor and new last element

			fact(f)
		}
	}
}

func getNextTwo(p int, y int) (int, int) {
	var np int
	var nl int

	if p == 1 {
		p = 2
	}

	if y%p == 0 {
		return p, y / p

	} else {
		for p = getNextPrime(p); y%p != 0; p = getNextPrime(p) {
			// Do nothing until a prime divisor is found
		}
		return p, y / p
	}

	for np = p; y%np != 0; np = getNextPrime(np) {
		nl = y / np
		fmt.Println("Next pair: " + strconv.Itoa(np) + " " + strconv.Itoa(nl))

		return np, nl

	}
	fmt.Println("Next pair: " + strconv.Itoa(p) + " " + strconv.Itoa(nl))

	return p, np
}

func getNextPrime(p int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137,
		139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223,
		227}

	// TODO: Develop a PANIC for primes out of range

	el := 0

	for i := 0; p >= primes[i]; i++ {
		el = primes[i+1]
	}
	return el

}

func isPrime(value int) bool { // Most basic implementation
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func fProduct(s []int) int {
	var i, n int
	for i, n = 0, s[0]; i < len(s); i++ {
		n = n * s[i]
	}
	return n
}

func commonFactors(s1 []int, s2 []int) []int {

	s := []int{}

	for i := 0; (len(s1) != 0) && (len(s2) != 0); i++ {
		// Test of match at the 0 element
		n1 := s1[0]
		n2 := s2[0]
		if n1 == n2 {
			s = append(s, n1)
			s1 = s1[1:]
			s2 = s2[1:]
		} else if n1 < n2 {
			s1 = s1[1:]

		} else {
			s2 = s2[1:]
		}
	}

	return s

}
