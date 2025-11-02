package linkedList

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()

	require.Equal(t, ll.len, 0)
}

func TestAddFirst(t *testing.T) {
	ll := NewLinkedList()

	ll.AddFirst("Hello,", "World!")
	require.Equal(t, ll.firstNode, ll.lastNode)
	
	for i := range 10 {
		ll.AddFirst(strconv.Itoa(i), strconv.Itoa(i))
		require.NotEqual(t, ll.firstNode, ll.lastNode)
		require.Equal(t, ll.firstNode.key, strconv.Itoa(i))
	}
}

func TestDeleteLast(t *testing.T) {
	ll := NewLinkedList()

	for i := range 10 {
		ll.AddFirst(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := range 9 {
		ll.DeleteLast()
		require.Equal(t, ll.lastNode.key, strconv.Itoa(i + 1))
	}
}

func TestDeleteElement(t *testing.T) {
	ll := NewLinkedList()

	for i := range 10 {
		ll.AddFirst(strconv.Itoa(i), strconv.Itoa(i))
	}

	for i := 0; i < 3; i++ {
		oldFirstKey := ll.firstNode.key
		ll.DeleteElement(ll.firstNode)
		require.NotEqual(t, oldFirstKey, ll.firstNode.key)
	}

	for i := 0; i < 3; i++ {
		oldLastKey := ll.lastNode.key
		ll.DeleteElement(ll.lastNode)
		require.NotEqual(t, oldLastKey, ll.lastNode.key)
	}

	for i := 0; i < 2; i++ {
		ll.DeleteElement(ll.firstNode.next)
	}

	require.Equal(t, ll.len, 2)
	require.Equal(t, ll.firstNode.key, "6")
	require.Equal(t, ll.lastNode.key, "3")
}

func TestDeleteLastFromOneElementList(t *testing.T) {
	ll := NewLinkedList()

	_ = ll.AddFirst("key", "value")

	ll.DeleteLast()

	require.Equal(t, ll.len, 0)
}

func TestDeleteFromOneElementList(t *testing.T) {
	ll := NewLinkedList()

	el := ll.AddFirst("key", "value")

	ll.DeleteElement(el)

	require.Equal(t, ll.len, 0)
}