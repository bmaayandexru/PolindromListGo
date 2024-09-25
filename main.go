package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value int
	Next  *Node
}

func CreatePoliList(n int) *Node {
	var (
		pprev *Node = nil
		p     *Node = nil
	)
	r := rand.New(rand.NewSource(99))
	s := make([]int, n)
	for i := 0; i < n/2+1; i++ {
		s[i] = r.Intn(10)
		s[len(s)-1-i] = s[i]
	}
	// вносим ошибку
	i := r.Intn(len(s))
	s[i] = 10 - s[i]
	fmt.Println(s)
	for _, v := range s {
		p = new(Node)
		p.Value = v
		p.Next = pprev
		pprev = p
	}
	PrintList(p)
	return p
}

func PrintList(h *Node) {
	cur := h
	fmt.Println("List:")
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}
}

func main() {
	var Head *Node
	fmt.Println("PolindromListGo...")
	Head = CreatePoliList(13)
	PrintList(Head)
}
