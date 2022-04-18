package main

import "fmt"

// pass slice as function argument
func addValue(foo []string) {
	foo = append(foo, "c")
	fmt.Println("modify foo", foo)
}

// todo: 这里真的不懂
func main() {
	foo := []string{"a", "b"}
	fmt.Println("before foo:", foo)
	addValue(foo)
	fmt.Println("after foo:", foo) // 没变
	fmt.Println("===============================")
	bar := foo[:1]
	fmt.Println("bar:", bar) //  [a]
	s1 := append(bar, "c")
	fmt.Println("bar:", bar)
	fmt.Println("foo:", foo) //  [a c]
	fmt.Println("s1:", s1)   //  [a c]
	fmt.Println("===============================")
	s2 := append(bar, "d")
	fmt.Println("bar:", bar)
	fmt.Println("foo:", foo) // [a d]
	fmt.Println("s2:", s2)   // [a d]
	fmt.Println("===============================")
	s3 := append(bar, "e", "f")
	fmt.Println("bar:", bar)
	fmt.Println("foo:", foo) //[a d]
	fmt.Println("s3:", s3)   //[a e f]
}
