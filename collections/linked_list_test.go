package collections

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestQueue(t *testing.T) {
	a := assert.New(t)

	q := NewLinkedList()
	a.Nil(q.head)
	a.Nil(q.tail)
	a.Empty(q.Contents())
	a.Nil(q.Dequeue())
	a.Equal(q.head, q.tail)
	a.Nil(q.Peek())
	a.Nil(q.PeekBack())
	a.Equal(0, q.Len())

	q.Enqueue("foo")
	a.NotNil(q.head)
	a.Nil(q.head.Previous)
	a.NotNil(q.tail)
	a.Equal(q.head, q.tail)
	a.Equal(1, q.Len())
	a.Equal("foo", q.Peek())
	a.Equal("foo", q.PeekBack())

	q.Enqueue("bar")
	a.NotNil(q.head)
	a.NotNil(q.head.Previous)
	a.Nil(q.head.Previous.Previous)
	a.Equal(q.head.Previous, q.tail)
	a.NotNil(q.tail)
	a.NotEqual(q.head, q.tail)
	a.Equal(2, q.Len())
	a.Equal("foo", q.Peek())
	a.Equal("bar", q.PeekBack())

	q.Enqueue("baz")
	a.NotNil(q.head)
	a.NotNil(q.head.Previous)
	a.NotNil(q.head.Previous.Previous)
	a.Nil(q.head.Previous.Previous.Previous)
	a.Equal(q.head.Previous.Previous, q.tail)
	a.NotNil(q.tail)
	a.NotEqual(q.head, q.tail)
	a.Equal(3, q.Len())
	a.Equal("foo", q.Peek())
	a.Equal("baz", q.PeekBack())

	q.Enqueue("fizz")
	a.NotNil(q.head)
	a.NotNil(q.head.Previous)
	a.NotNil(q.head.Previous.Previous)
	a.NotNil(q.head.Previous.Previous.Previous)
	a.Nil(q.head.Previous.Previous.Previous.Previous)
	a.Equal(q.head.Previous.Previous.Previous, q.tail)
	a.NotNil(q.tail)
	a.NotEqual(q.head, q.tail)
	a.Equal(4, q.Len())
	a.Equal("foo", q.Peek())
	a.Equal("fizz", q.PeekBack())

	values := q.Contents()
	a.Len(4, values)
	a.Equal("foo", values[0])
	a.Equal("bar", values[1])
	a.Equal("baz", values[2])
	a.Equal("fizz", values[3])

	shouldBeFoo := q.Dequeue()
	a.Equal("foo", shouldBeFoo)
	a.NotNil(q.head)
	a.NotNil(q.head.Previous)
	a.NotNil(q.head.Previous.Previous)
	a.Nil(q.head.Previous.Previous.Previous)
	a.Equal(q.head.Previous.Previous, q.tail)
	a.NotNil(q.tail)
	a.NotEqual(q.head, q.tail)
	a.Equal(3, q.Len())
	a.Equal("bar", q.Peek())
	a.Equal("fizz", q.PeekBack())

	shouldBeBar := q.Dequeue()
	a.Equal("bar", shouldBeBar)
	a.NotNil(q.head)
	a.NotNil(q.head.Previous)
	a.Nil(q.head.Previous.Previous)
	a.Equal(q.head.Previous, q.tail)
	a.NotNil(q.tail)
	a.NotEqual(q.head, q.tail)
	a.Equal(2, q.Len())
	a.Equal("baz", q.Peek())
	a.Equal("fizz", q.PeekBack())

	shouldBeBaz := q.Dequeue()
	a.Equal("baz", shouldBeBaz)
	a.NotNil(q.head)
	a.Nil(q.head.Previous)
	a.NotNil(q.tail)
	a.Equal(q.head, q.tail)
	a.Equal(1, q.Len())
	a.Equal("fizz", q.Peek())
	a.Equal("fizz", q.PeekBack())

	shouldBeFizz := q.Dequeue()
	a.Equal("fizz", shouldBeFizz)
	a.Nil(q.head)
	a.Nil(q.tail)
	a.Nil(q.Peek())
	a.Nil(q.PeekBack())
	a.Equal(0, q.Len())
}
