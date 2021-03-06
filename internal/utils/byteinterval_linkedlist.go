package utils

type ByteIntervalElement struct {
	next, prev *ByteIntervalElement
	list       *ByteIntervalList
	Value      ByteInterval
}

func (e *ByteIntervalElement) Next() *ByteIntervalElement {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}
func (e *ByteIntervalElement) Prev() *ByteIntervalElement {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// ByteIntervalList is a linked list of ByteIntervals.
type ByteIntervalList struct {
	root ByteIntervalElement // sentinel list element, only &root, root.prev, and root.next are used
	len  int                 // current list length excluding (this) sentinel element
}

// Init initializes or clears list l.
func (l *ByteIntervalList) Init() *ByteIntervalList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// NewByteIntervalList returns an initialized list.
func NewByteIntervalList() *ByteIntervalList { return new(ByteIntervalList).Init() }

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *ByteIntervalList) Len() int { return l.len }

// Front returns the first element of list l or nil if the list is empty.
func (l *ByteIntervalList) Front() *ByteIntervalElement {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *ByteIntervalList) Back() *ByteIntervalElement {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// lazyInit lazily initializes a zero List value.
func (l *ByteIntervalList) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (l *ByteIntervalList) insert(e, at *ByteIntervalElement) *ByteIntervalElement {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *ByteIntervalList) insertValue(v ByteInterval, at *ByteIntervalElement) *ByteIntervalElement {
	return l.insert(&ByteIntervalElement{Value: v}, at)
}

// remove removes e from its list, decrements l.len, and returns e.
func (l *ByteIntervalList) remove(e *ByteIntervalElement) *ByteIntervalElement {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *ByteIntervalList) Remove(e *ByteIntervalElement) ByteInterval {
	if e.list == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *ByteIntervalList) PushFront(v ByteInterval) *ByteIntervalElement {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *ByteIntervalList) PushBack(v ByteInterval) *ByteIntervalElement {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *ByteIntervalList) InsertBefore(v ByteInterval, mark *ByteIntervalElement) *ByteIntervalElement {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *ByteIntervalList) InsertAfter(v ByteInterval, mark *ByteIntervalElement) *ByteIntervalElement {
	if mark.list != l {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ByteIntervalList) MoveToFront(e *ByteIntervalElement) {
	if e.list != l || l.root.next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *ByteIntervalList) MoveToBack(e *ByteIntervalElement) {
	if e.list != l || l.root.prev == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.insert(l.remove(e), l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ByteIntervalList) MoveBefore(e, mark *ByteIntervalElement) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *ByteIntervalList) MoveAfter(e, mark *ByteIntervalElement) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.insert(l.remove(e), mark)
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ByteIntervalList) PushBackList(other *ByteIntervalList) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *ByteIntervalList) PushFrontList(other *ByteIntervalList) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
