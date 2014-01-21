package list

import (
	"errors"
)

// Node is a node of the list
type Node struct {
	next  *Node       // The node after this node in the list
	list  *LinkedList // The list to which this element belongs
	Value interface{} // The value stored with this node
}

// Next returns the next node or nil
func (n *Node) Next() *Node {
	if i := n.next; n.list != nil {
		return i
	}

	return nil
}

// LinkedList is a single linked list
type LinkedList struct {
	first *Node // The first node of the list
	last  *Node // The last node of the list
	len   int   // The current list length
}

// New returns an initialized list
func New() *LinkedList {
	return new(LinkedList).Init()
}

// Init initializes or clears the list
func (l *LinkedList) Init() *LinkedList {
	l.Clear()

	return l
}

// Clear removes all nodes from the list
func (l *LinkedList) Clear() {
	i := l.first

	for i != nil {
		j := i.Next()

		i.list = nil
		i.next = nil

		i = j
	}

	l.first = nil
	l.last = nil
	l.len = 0
}

// Len returns the curren list length
func (l *LinkedList) Len() int {
	return l.len
}

// First returns the first node of the list or nil
func (l *LinkedList) First() *Node {
	return l.first
}

// Last returns the last node of the list or nil
func (l *LinkedList) Last() *Node {
	return l.last
}

// Get returns the node with the given index or nil
func (l *LinkedList) Get(i int) (*Node, error) {
	if i < 0 || i >= l.len {
		return nil, errors.New("index bouunds out of range")
	}

	j := 0

	for n := l.First(); n != nil; n = n.Next() {
		if i == j {
			return n, nil
		}

		j++
	}

	return nil, nil
}

// Set replaces the value in the list with the given value
func (l *LinkedList) Set(i int, v interface{}) error {
	if i < 0 || i >= l.len {
		return errors.New("index bouunds out of range")
	}

	j := 0

	for n := l.First(); n != nil; n = n.Next() {
		if i == j {
			n.Value = v

			return nil
		}

		j++
	}

	return nil
}

// Copy returns an exact copy of the list
func (l *LinkedList) Copy() *LinkedList {
	n := New()

	for i := l.First(); i != nil; i = i.Next() {
		n.Push(i.Value)
	}

	return n
}

// newNode initializes a new node for the list
func (l *LinkedList) newNode(v interface{}) *Node {
	return &Node{
		list:  l,
		Value: v,
	}
}

// findParent returns the parent to a given node or nil
func (l *LinkedList) findParent(c *Node) *Node {
	if c == nil || c.list != l {
		return nil
	}

	var p *Node

	for i := l.First(); i != nil; i = i.Next() {
		if i == c {
			return p
		}

		p = i
	}

	return nil
}

// insertAfter creates a new node from a value, inserts it after a given node and returns the new one
func (l *LinkedList) insertAfter(v interface{}, p *Node) *Node {
	if p != nil && p.list != l {
		return nil
	}

	n := l.newNode(v)

	// insert first node
	if p == nil {
		l.first = n
		l.last = n
	} else {
		p.next = n

		if p == l.last {
			l.last = n
		}
	}

	l.len++

	return n
}

// insertBefore creates a new node from a value, inserts it before a given node and returns the new one
func (l *LinkedList) insertBefore(v interface{}, p *Node) *Node {
	if p != nil && p.list != l {
		return nil
	}

	n := l.newNode(v)

	// insert first node
	if p == nil {
		l.first = n
		l.last = n
	} else {
		if p == l.first {
			l.first = n
		} else {
			pp := l.findParent(p)

			pp.next = n
		}

		n.next = p
	}

	l.len++

	return n
}

// Remove removes a given node from the list
func (l *LinkedList) Remove(c *Node) *Node {
	if c == nil || c.list != l || l.len == 0 {
		return nil
	}

	r := c

	if c == l.first {
		l.first = c.next

		// c is the last node
		if c == l.last {
			l.last = nil
		}
	} else {
		p := l.findParent(c)

		p.next = c.next

		if c == l.last {
			l.last = p
		}
	}

	r.list = nil
	r.next = nil

	l.len--

	return r
}

// Pop removes and returns the last node or nil
func (l *LinkedList) Pop() *Node {
	return l.Remove(l.last)
}

// Push creates a new node from a value, inserts it as the last node and returns it
func (l *LinkedList) Push(v interface{}) *Node {
	return l.insertAfter(v, l.last)
}

// Shift removes and returns the first node or nil
func (l *LinkedList) Shift() *Node {
	return l.Remove(l.first)
}

// Unshift creates a new node from a value, inserts it as the first node and returns it
func (l *LinkedList) Unshift(v interface{}) *Node {
	return l.insertBefore(v, l.first)
}

// Contains returns true if the value exists in the list
func (l *LinkedList) Contains(v interface{}) bool {
	_, ok := l.IndexOf(v)

	return ok
}

// IndexOf returns the index of an occurence of the given value and true or -1 and false if the value does not exist
func (l *LinkedList) IndexOf(v interface{}) (int, bool) {
	i := 0

	for n := l.First(); n != nil; n = n.Next() {
		if n.Value == v {
			return i, true
		}

		i++
	}

	return -1, false
}