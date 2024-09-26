package main

import (
	"fmt"
	"math/rand"
	"time"
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
	if n == 0 {
		return p
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]int, n)
	for i := 0; i < n/2+1; i++ {
		s[i] = r.Intn(10)
		s[len(s)-1-i] = s[i]
	}
	// вносим ошибку
	//	i := r.Intn(len(s))
	//	s[i] = 10 - s[i]
	fmt.Println(s)
	for _, v := range s {
		p = new(Node)
		p.Value = v
		p.Next = pprev
		pprev = p
	}
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

func IsPolindrom(h *Node) bool {
	var (
		count int   = 0
		pend  *Node = nil // указательна конец списка
		hr    *Node = nil
		hl    *Node = nil
		sref  *Node = nil
	)
	if h == nil {
		// список пуст. ничего не делаем и сразу выходим
		return false
	}
	pend = h   // сделали шаг
	count += 1 // инкркмент счетчика
	fmt.Println(pend, count)
	for pend.Next != nil {
		pend = pend.Next // делаем шаг
		count += 1       // и инкремент
		fmt.Println(pend, count)
		if count < 3 { // count == 2
			// инициализация hr и hl
			hr = pend
			hl = h
		} else {
			// переворачивание левой половины в обратном направлении
			if count%2 == 0 {
				// count четный. сдвигаем hl и hr на один элемент вправо
				// и разворачиваем ссылку между стратовыми hl  hr
				sref = hl
				hl = hl.Next
				hr = hr.Next
				hl.Next = sref // разворот
			}

		}
	}
	if count == 1 {
		// список из одного элемента считаем полиндромом
		return true
	}
	// тут начало сравнения hr и hl
	for {
		if hr.Value == hl.Value {
			// следующий шаг
		} else {
			// не полиндром
			// доперевернуть левую половину и выйти с false
			return false
		}
		if hl == h {
			// условие окончания
		}
	}
	return false
}
func main() {
	var Head *Node
	fmt.Println("PolindromListGo...")
	Head = CreatePoliList(7)
	PrintList(Head)
	fmt.Println("Список полиндром: ", IsPolindrom(Head))
}
