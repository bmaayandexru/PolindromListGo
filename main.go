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

// CreatePoliList() создает список-полиндром
// если seterr == true случайно портит один из элементов, чтобы список не был полиндромом
func CreatePoliList(n int, seterr bool) *Node {
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
	if seterr {
		i := r.Intn(len(s))
		if i == len(s)/2 {
			i--
		}
		s[i] = 10 - s[i]
	}
	fmt.Println(s)
	for _, v := range s {
		p = new(Node)
		p.Value = v
		p.Next = pprev
		pprev = p
	}
	return p
}

// PrintList выводит на печать список
func PrintList(h *Node) {
	cur := h
	fmt.Println("List:")
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}
}

// IsPolindrom() возвращает true если списо является полинромом и false если нет
func IsPolindrom(h *Node) bool {
	var (
		count int   = 0
		pend  *Node = nil // указательна конец списка
		hr    *Node = nil // указатель на голову правого списка
		hl    *Node = nil // указатель на голову левого списка
		s_hl  *Node = nil // переменная для сохранения головы левого списка
		s_hr  *Node = nil // переменная для сохранения головы правого списка
	)
	if h == nil {
		// список пуст. ничего не делаем и сразу выходим
		return false
	}
	pend = h   // сделали шаг
	count += 1 // инкремент счетчика
	fmt.Println(pend, count)
	for pend.Next != nil {
		pend = pend.Next // делаем шаг
		count += 1       // и инкремент
		if count < 3 {   // count == 2
			// инициализация hr и hl
			hr = pend
			hl = h
		} else {
			// переворачивание левой половины в обратном направлении
			if count%2 == 0 {
				// count четный. сдвигаем hl и hr на один элемент вправо
				// и разворачиваем ссылку между стратовыми hl  hr
				s_hr = hr
				s_hl = hl
				hr = hr.Next
				hl = s_hr
				hl.Next = s_hl
			}
		}
	}
	if count == 1 {
		// список из одного элемента считаем полиндромом
		return true
	}
	// тут начало сравнения hr и hl
	// сохраняем ссылки на головы списков для возврата в правый список
	s_hr = hr
	s_hl = hl
	// сдвиг hr на один элемент если count нечетный
	if count%2 == 1 {
		hr = hr.Next
	}
	for {
		if hr.Value == hl.Value {
			if hl == h { // или hr == pe
				// условие окончания
				// *** возврат списка в исходное положение
				hl = s_hl
				hr = s_hr
				for hl != h {
					s_hl = hl.Next
					hl.Next = hr
					hr = hl
					hl = s_hl
				}
				h.Next = hr
				return true
			}
			// следующий шаг hl(hr) = hl(hr).Next
			// движение по левому и правому списку
			hl = hl.Next
			hr = hr.Next
		} else {
			// не полиндром
			// *** возврат списка в исходное положение
			hl = s_hl
			hr = s_hr
			for hl != h {
				s_hl = hl.Next
				hl.Next = hr
				hr = hl
				hl = s_hl
			}
			h.Next = hr
			return false
		}
	}
}
func main() {
	var Head *Node
	fmt.Println("PolindromListGo...")
	Head = CreatePoliList(8, false)
	fmt.Println("Список полиндром: ", IsPolindrom(Head))
	PrintList(Head)
	Head = CreatePoliList(8, true)
	fmt.Println("Список полиндром: ", IsPolindrom(Head))
	PrintList(Head)
	Head = CreatePoliList(7, true)
	PrintList(Head)
	fmt.Println("Список полиндром: ", IsPolindrom(Head))
}
