package linkedList

type LinkedList struct {
	firstNode *Element
	lastNode  *Element
	len       int
}

type Element struct {
	key      string
	value    string
	previous *Element
	next     *Element
}

func NewLinkedList() *LinkedList {
	return &LinkedList{len: 0}
}

func (ll *LinkedList) AddFirst(key, value string) *Element {
	newElement := &Element{key: key, value: value}

	return ll.addFirst(newElement)
}

func (ll *LinkedList) addFirst(newElement *Element) *Element {
	if ll.len == 0 {
		ll.firstNode = newElement
		ll.lastNode = newElement
		ll.len++
		return newElement
	}

	ll.firstNode.previous = newElement
	newElement.next = ll.firstNode
	ll.firstNode = newElement
	ll.len++

	return newElement
}

func (ll *LinkedList) DeleteLast() *Element {
	switch ll.len {
	case 0:
		return nil

	case 1:
		el := ll.lastNode
		ll.firstNode = nil
		ll.lastNode = nil
		ll.len--
		return el

	default:
		el := ll.lastNode
		newLastNode := ll.lastNode.previous
		ll.lastNode.previous = nil
		newLastNode.next = nil
		ll.lastNode = newLastNode
		ll.len--
		return el
	}
}

func (ll *LinkedList) MoveToFirst(newFirstElement *Element) {
	switch newFirstElement {
	case ll.firstNode:
		return

	default:
		ll.Delete(newFirstElement)
		ll.addFirst(newFirstElement)
	}
}

func (ll *LinkedList) Delete(elementToDelete *Element) *Element {
	switch {
	case elementToDelete == ll.firstNode && elementToDelete == ll.lastNode:
		ll.firstNode = nil
		ll.lastNode = nil
		ll.len--
		return elementToDelete

	case elementToDelete == ll.firstNode:
		newFirstNode := elementToDelete.next
		newFirstNode.previous = nil
		elementToDelete.next = nil
		ll.firstNode = newFirstNode
		ll.len--
		return elementToDelete

	case elementToDelete == ll.lastNode:
		return ll.DeleteLast()

	default:
		previous := elementToDelete.previous
		next := elementToDelete.next
		elementToDelete.previous = nil
		elementToDelete.next = nil
		previous.next = next
		next.previous = previous
		ll.len--
		return elementToDelete
	}
}

func (ll *LinkedList) GetLen() int {
	return ll.len
}

func (el *Element) GetValue() string {
	return el.value
}

func (el *Element) GetKey() string {
	return el.key
}