package main

import "fmt"

// Generate natural seri number: 2,3,4,...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter: delete the number which is divisible by a prime number to find prime number
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch

		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)

	}
}

/**
This is a pretty convoluted example. In both functions, go func(){...}() creates an anonymous goroutine and runs it asynchronously, then returns the channel which will receive values from the goroutine. PrimeFilter returns a channel which will receive numbers not divisible by a certain candidate.

The idea is that prime := <-ch always takes the first element from the channel. So, to visualize the flow:

GenerateNatural() starts by sending numbers 2, 3, 4... to ch.

First loop iteration:

a. prime := <-ch reads the first (prime) number 2.

b. PrimeFilter(ch, 2) then continues receiving the rest of the numbers (3, 4, 5, ...), and sends numbers not divisible by 2 to the output channel. So, channel returned by PrimeFilter(ch, 2) will receive numbers (3, 5, 7, ...).

c. ch = PrimeFilter(ch, prime) in the main function now replaces the local ch variable with the output of PrimeFilter(ch, 2) from the previous step.

Second loop iteration:

a. prime := <-ch reads the first (prime) number from the current ch instance (this first number is 3).

b. PrimeFilter(ch, 3) then continues receiving the (already filtered) numbers, except for the first one (so, 5, 7, 9, ...), and sends numbers not divisible by 3 to the output channel. So, channel returned by PrimeFilter(ch, 2) will receive numbers 5, 7, 11, ..., because 9 is divisible by 3.

c. ch = PrimeFilter(ch, prime) in the main function now replaces the local ch variable with the output of PrimeFilter(ch, 3) from the previous step.

...
*/
