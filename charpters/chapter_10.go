package main

import "fmt"

// map 无序map

func main() {
	m := make(map[string]int)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%d", i)
		m[s] = i
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m)

	for k := range m {
		fmt.Println(k)
		delete(m, k)
	}
	m["test"] = 1
	fmt.Println(m)
}
