package unrolledlinkedlist

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

var v = []interface{}{1, "a", 2, "b", 3, "c", 4, "d"}
var vLen = len(v)

func fillList(t *testing.T, l *UnrolledLinkedList) {
	for i, va := range v {
		l.Push(va)

		Equal(t, l.Len(), i+1)
		Equal(t, l.First().Value(), v[0])
		Equal(t, l.Last().Value(), va)
	}

	Equal(t, l.Len(), vLen)
}

func newFilledList(t *testing.T) *UnrolledLinkedList {
	l := New(4)

	fillList(t, l)

	return l
}

func printList(l *UnrolledLinkedList) {
	fmt.Printf("Len: %d\n", l.len)
	for n := l.first; n != nil; n = n.Next() {
		fmt.Printf("\t%+v\n", n.values)
	}
}

func TestBasic(t *testing.T) {
	l := New(4)

	Equal(t, l.Len(), 0)
	Nil(t, l.First())
	Nil(t, l.Last())
	n, ok := l.Pop()
	Nil(t, n)
	False(t, ok)
	n, ok = l.Shift()
	Nil(t, n)
	False(t, ok)

	fillList(t, l)

	i := 0
	c := l.First()

	for i < vLen {
		Equal(t, v[i], c.Value())

		i++
		if i < vLen {
			True(t, c.Next())
		}
	}

	False(t, c.Next())
	Nil(t, c.Value())

	i = vLen - 1
	n, ok = l.Pop()

	for i > -1 && n != nil {
		Equal(t, v[i], n)
		True(t, ok)
		Equal(t, l.Len(), i)

		i--
		n, ok = l.Pop()
	}

	Equal(t, i, -1)
	Nil(t, n)
	False(t, ok)
	Equal(t, l.Len(), 0)

	for i, va := range v {
		l.Unshift(va)

		Equal(t, l.Len(), i+1)
		Equal(t, l.First().Value(), va)
		Equal(t, l.Last().Value(), v[0])
	}

	Equal(t, l.Len(), vLen)

	i = vLen - 1
	n, ok = l.Shift()

	for i > -1 && n != nil {
		Equal(t, v[i], n)
		True(t, ok)
		Equal(t, l.Len(), i)

		i--
		n, ok = l.Shift()
	}

	Equal(t, i, -1)
	Nil(t, n)
	Equal(t, l.Len(), 0)
}

/*

func TestToArray(t *testing.T) {
	l := New()
	Equal(t, l.ToArray(), []interface{}{})

	fillList(t, l)
	Equal(t, l.ToArray(), v)

	l.Shift()
	Equal(t, l.ToArray(), v[1:])

	l.Pop()
	Equal(t, l.ToArray(), v[1:3])
}

func TestRemove(t *testing.T) {
	l1 = newFilledList(t)

	// out of bound
	_, err := l1.RemoveAt(-1)
	NotNil(t, err)
	_, err = l1.RemoveAt(l1.Len())
	NotNil(t, err)

	// Remove Middle
	n1, err = l1.RemoveAt(1)
	Nil(t, err)
	Equal(t, n1.Value, v[1])
	Nil(t, n1.Next())
	Nil(t, n1.Previous())
	n2, _ = l1.Get(1)
	Equal(t, n2.Value, v[2])
	Equal(t, n2.Next().Value, v[3])
	Equal(t, l1.Len(), 3)

	// Remove First
	n1, err = l1.RemoveAt(0)
	Nil(t, err)
	Equal(t, n1.Value, v[0])
	Nil(t, n1.Next())
	n2 = l1.First()
	Equal(t, n2.Value, v[2])
	Equal(t, n2.Next().Value, v[3])
	Equal(t, l1.Len(), 2)

	// Remove Last
	n1, err = l1.RemoveAt(1)
	Nil(t, err)
	Equal(t, n1.Value, v[3])
	Nil(t, n1.Next())
	n2 = l1.Last()
	Equal(t, n2.Value, v[2])
	Nil(t, n2.Next())
	Equal(t, l1.Len(), 1)

	// Remove very last node
	n1, err = l1.RemoveAt(0)
	Nil(t, err)
	Equal(t, n1.Value, v[2])
	Nil(t, n1.Next())
	Equal(t, l1.Len(), 0)
	Nil(t, l1.First())
	Nil(t, l1.Last())
}

func TestRemoveOccurrence(t *testing.T) {
	l := New()

	for i := 0; i < 5; i++ {
		l.Push(i % 2)
	}

	n := l.RemoveFirstOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 4)
	Equal(t, l.First().Value, 1)
	Equal(t, l.First().Next().Value, 0)
	Equal(t, l.First().Next().Next().Value, 1)
	Equal(t, l.Last().Value, 0)

	n = l.RemoveFirstOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 3)
	Equal(t, l.First().Value, 1)
	Equal(t, l.First().Next().Value, 1)
	Equal(t, l.Last().Value, 0)

	n = l.RemoveFirstOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 2)
	Equal(t, l.First().Value, 1)
	Equal(t, l.Last().Value, 1)

	n = l.RemoveFirstOccurrence(0)
	Nil(t, n)
	Equal(t, l.Len(), 2)
	Equal(t, l.First().Value, 1)
	Equal(t, l.Last().Value, 1)

	l.Clear()

	for i := 0; i < 5; i++ {
		l.Push(i % 2)
	}

	n = l.RemoveLastOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 4)
	Equal(t, l.First().Value, 0)
	Equal(t, l.First().Next().Value, 1)
	Equal(t, l.First().Next().Next().Value, 0)
	Equal(t, l.Last().Value, 1)

	n = l.RemoveLastOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 3)
	Equal(t, l.First().Value, 0)
	Equal(t, l.First().Next().Value, 1)
	Equal(t, l.Last().Value, 1)

	n = l.RemoveLastOccurrence(0)
	Equal(t, n.Value, 0)
	Nil(t, n.Next())
	Equal(t, l.Len(), 2)
	Equal(t, l.First().Value, 1)
	Equal(t, l.Last().Value, 1)

	n = l.RemoveLastOccurrence(0)
	Nil(t, n)
	Equal(t, l.Len(), 2)
	Equal(t, l.First().Value, 1)
	Equal(t, l.Last().Value, 1)
}

func TestClear(t *testing.T) {
	l := newFilledList(t)

	l.Clear()

	Equal(t, l.Len(), 0)
	Nil(t, l.First())
	Nil(t, l.Last())
	Nil(t, l.Pop())
}

func TestCopy(t *testing.T) {
	l1 := newFilledList(t)

	l2 := l1.Copy()

	Equal(t, l1.Len(), l2.Len())

	n1 := l1.First()
	n2 := l2.First()

	for n1 != nil && n2 != nil {
		Equal(t, n1.Value, n2.Value)

		n1 = n1.Next()
		n2 = n2.Next()
	}

	Nil(t, n1)
	Nil(t, n2)
}

*/

func TestFind(t *testing.T) {
	l := New(4)

	for _, vi := range v {
		f, ok := l.IndexOf(vi)
		Equal(t, f, -1)
		Equal(t, ok, false)

		ok = l.Contains(vi)
		Equal(t, ok, false)
	}

	fillList(t, l)

	for i, vi := range v {
		f, ok := l.IndexOf(vi)
		Equal(t, f, i)
		Equal(t, ok, true)

		ok = l.Contains(vi)
		Equal(t, ok, true)
	}

	l.Clear()

	f, ok := l.IndexOf(0)
	Equal(t, f, -1)
	Equal(t, ok, false)

	f, ok = l.LastIndexOf(0)
	Equal(t, f, -1)
	Equal(t, ok, false)

	for i := 0; i < 4; i++ {
		l.Push(0)

		f, ok = l.IndexOf(0)
		Equal(t, f, 0)
		Equal(t, ok, true)

		f, ok = l.LastIndexOf(0)
		Equal(t, f, i)
		Equal(t, ok, true)
	}

	// not found in nonempty list
	f, ok = l.IndexOf(100)
	Equal(t, f, -1)
	Equal(t, ok, false)

	f, ok = l.LastIndexOf(100)
	Equal(t, f, -1)
	Equal(t, ok, false)
}

/*

func TestGetSet(t *testing.T) {
	l := New()

	for i := range v {
		n, err := l.Get(i)

		Nil(t, n)
		NotNil(t, err)

		err = l.Set(i, i+10)

		NotNil(t, err)

		n, err = l.Get(i)

		Nil(t, n)
		NotNil(t, err)
	}

	fillList(t, l)

	for i, vi := range v {
		n, err := l.Get(i)

		Equal(t, n.Value, vi)
		Nil(t, err)

		err = l.Set(i, i+10)

		Nil(t, err)

		n, err = l.Get(i)

		Equal(t, n.Value, i+10)
		Nil(t, err)
	}
}

func TestAddLists(t *testing.T) {
	l1 := New()
	l1.Push(3)
	l1.Push(4)

	l2 := New()
	l2.Push(5)
	l2.Push(6)

	l3 := New()
	l3.Push(2)
	l3.Push(1)

	l1.PushList(l2)
	Equal(t, l1.ToArray(), []interface{}{3, 4, 5, 6})

	l1.UnshiftList(l3)
	Equal(t, l1.ToArray(), []interface{}{1, 2, 3, 4, 5, 6})
}

func TestFunc(t *testing.T) {
	l := newFilledList(t)

	Equal(t, v[1], l.GetFunc(func(n *Node) bool {
		return n.Value == "a"
	}).Value)
	Nil(t, l.GetFunc(func(n *Node) bool {
		return n.Value == "z"
	}))

	l.SetFunc(func(n *Node) bool {
		return n.Value == 2
	}, 3)
	Equal(t, l.ToArray(), []interface{}{1, "a", 3, "b"})
	l.SetFunc(func(n *Node) bool {
		return n.Value == "z"
	}, 4)
	Equal(t, l.ToArray(), []interface{}{1, "a", 3, "b"})
}

*/
