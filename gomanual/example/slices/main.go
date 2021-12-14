package main

import "fmt"

func main() {
	var a = []string{"linux", "linguagem", "concorrente", "2021"}
	var b = []string{"quero", " aprender ", "Go"}

	// println("append")
	// a = append(a, b...)
	// fmt.Println(a)
	// fmt.Println(b)

	// b = append(a, a...)
	// fmt.Println(b)

	println(".........copy")
	b = make([]string, len(a))
	copy(b, a)
	fmt.Println(b)

	// b = append(a[:0:0], a...)
	b = append(make([]string, 0, len(a)), a...)
	//copy(b, a)
	fmt.Println(b)
	fmt.Println(a)

	println("recortar")
	// recortar - Cut
	i := 0
	j := 2
	a = append(a[:i], a[j:]...)
	fmt.Println(a)

	println("delete")
	fmt.Println(b)
	i = 1
	//b = append(b[:i], b[i+1:]...)
	// or
	b = b[:i+copy(a[i:], a[i+1:])]
	fmt.Println(b)

	// push
	println("push")
	var x = []int{10, 30, 90}
	var a1 []int
	a1 = append(a1, x...)
	fmt.Println(a1)

	println("pop")
	// x = a1[len(a1)-1]
	// a1 = a1[:len(a1)-1]
	// fmt.Println(x, a1)

}
