package linkedList

type LinkedList struct {
	firstNode *Element
	lastNode *Element
	len int
}

type Element struct {
	key string
	value string
	previous *Element
	next *Element
}

func NewLinkedList() *LinkedList {
	return &LinkedList{len: 0}
}

func (ll *LinkedList) AddFirst(key, value string) *Element {
	newElement := Element{key: key, value: value}

	if ll.len == 0 {
		ll.firstNode = &newElement
		ll.lastNode = &newElement
		ll.len++
		return &newElement
	}

	ll.firstNode.previous = &newElement
	newElement.next = ll.firstNode
	ll.firstNode = &newElement
	ll.len++

	return &newElement
}

func (ll *LinkedList) DeleteLast() {
	switch ll.len {
	case 0:
		return

	case 1:
		ll.firstNode = nil
		ll.lastNode = nil
		ll.len--

	default:
		newLastNode := ll.lastNode.previous
		ll.lastNode.previous = nil
		newLastNode.next = nil
		ll.lastNode = newLastNode
		ll.len--
	}
}

func (ll *LinkedList) MoveToFirst(newFirstElement *Element) {
	switch newFirstElement {
	case ll.firstNode:
		return
	
	case ll.lastNode:
		newLastElement := newFirstElement.previous
		newFirstElement.previous = nil
		newLastElement.next = nil
		ll.lastNode = newLastElement
		fallthrough
	
	default:
		ll.firstNode.previous = newFirstElement
		newFirstElement.next = ll.firstNode
		ll.firstNode = newFirstElement
	}
}

func (ll *LinkedList) DeleteElement(elementToDelete *Element) {
	switch {
	case elementToDelete == ll.firstNode && elementToDelete == ll.lastNode:
		ll.firstNode = nil
		ll.lastNode = nil

	case elementToDelete == ll.firstNode:
		newFirstNode := elementToDelete.next
		newFirstNode.previous = nil
		elementToDelete.next = nil
		ll.firstNode = newFirstNode
	
	case elementToDelete == ll.lastNode:
		newLastElement := elementToDelete.previous
		newLastElement.next = nil
		elementToDelete.previous = nil
		ll.lastNode = newLastElement
	
	default:
		previous := elementToDelete.previous
		next := elementToDelete.next
		elementToDelete.previous = nil
		elementToDelete.next = nil
		previous.next = next
		next.previous = previous
	}
	
	ll.len--
}

func (ll *LinkedList) GetLastKey() string {
	return ll.lastNode.key
}

func (ll *LinkedList) GetLen() int {
	return ll.len
}

func (el *Element) GetValue() string {
	return el.value
}