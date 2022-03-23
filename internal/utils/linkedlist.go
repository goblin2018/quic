package utils

import "github.com/cheekybits/genny/generic"

type Item generic.Type

type ItemElement struct {
	next, prev *ItemElement
	list       *ItemList
	Value      Item
}

func (e *ItemElement) Next() *ItemElement {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type ItemList struct {
	root ItemElement
	len  int
}

func (e *ItemElement) Prev() *ItemElement {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (l *ItemList) Init() *ItemList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func NewItemList() *ItemList { return new(ItemList).Init() }

func (l *ItemList) Len() int {
	return l.len
}

func (l *ItemList) Front() *ItemElement {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}
func (l *ItemList) Back() *ItemElement {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}
func (l *ItemList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
func (l *ItemList) insert(e, at *ItemElement) *ItemElement {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}
func (l *ItemList) insertValue(v Item, at *ItemElement) *ItemElement {
	return l.insert(&ItemElement{Value: v}, at)
}
func (l *ItemList) remove(e *ItemElement) *ItemElement {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}
func (l *ItemList) Remove(e *ItemElement) Item {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}
func (l *ItemList) PushFront(v Item) *ItemElement {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}
func (l *ItemList) PushBack(v Item) *ItemElement {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}
func (l *ItemList) InsertBefore(v Item, mark *ItemElement) *ItemElement {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *ItemList) InsertAfter(v Item, mark *ItemElement) *ItemElement {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ItemList) MoveToFront(e *ItemElement) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ItemList) MoveToBack(e *ItemElement) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ItemList) MoveBefore(e, mark *ItemElement) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ItemList) MoveAfter(e, mark *ItemElement) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark)
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ItemList) PushBackList(other *ItemList) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ItemList) PushFrontList(other *ItemList) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
