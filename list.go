package doublylinked

// ListInerfase interface
type ListInerfase interface {
	Len()                    // длинна списка
	First()                  // первый Item
	Last()                   // последний Item
	PushFront(v interface{}) // добавить значение в начало
	PushBack(v interface{})  // добавить значение в конец
}

// List struct of the doubly linked list
type List struct {
	head  *Item
	tail  *Item
	count int
}

// First get first element returning Item
func (l *List) First() *Item {
	return l.head
}

// Last get last element returning Item
func (l *List) Last() *Item {
	return l.tail
}

// Len get last element returning count elements
func (l *List) Len() int {
	return l.count
}

// PushBack push back an item to the list
func (l *List) PushBack(v interface{}) {
	l.count++
	item := &Item{v, nil, nil, l}
	if l.head == nil {
		l.head = item
	} else {
		l.tail.next = item
		item.prev = l.tail
	}
	l.tail = item
}

// PushFront push an item forvard to the list
func (l *List) PushFront(v interface{}) {
	l.count++
	item := &Item{v, nil, nil, l}
	if l.tail == nil {
		l.tail = item
	} else {
		l.head.prev = item
		item.next = l.head
	}
	l.head = item
}

// ItemInerfase interface
type ItemInerfase interface {
	Value() interface{} // возвращает значение
	Next() *Item        // следующий Item
	Prev() *Item        // предыдущий
	Remove()            // удалить Item из списка
}

// Item item of the doubly linked list
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
	*List
}

// Next next iteration returning Item
func (i *Item) Next() *Item {
	return i.next
}

// Prev prev iteration returning Item
func (i *Item) Prev() *Item {
	return i.prev
}

// Value  returning interface
func (i *Item) Value() interface{} {
	return i.value
}

// Remove item from list
func (i *Item) Remove() {
	prevItem := i.Prev()
	nextItem := i.Next()
	if prevItem != nil {
		prevItem.next = nextItem
	} else {
		i.head = nextItem
	}
	if nextItem != nil {
		nextItem.prev = prevItem
	} else {
		i.tail = prevItem
	}
	i.count--
}
