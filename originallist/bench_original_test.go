// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package originallist

import (
	"container/list"
	"testing"
)

func BenchmarkList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := list.New()

		// Single element list
		e := l.PushFront(42)
		l.MoveToFront(e)
		l.MoveToBack(e)
		l.Remove(e)

		// Bigger list
		e2 := l.PushFront(2)
		e1 := l.PushFront(1)
		e3 := l.PushBack(3)
		e4 := l.PushBack(84)

		l.Remove(e2)

		l.MoveToFront(e3) // move from middle

		l.MoveToFront(e1)
		l.MoveToBack(e3) // move from middle

		l.MoveToFront(e3) // move from back
		l.MoveToFront(e3) // should be no-op

		l.MoveToBack(e3) // move from front
		l.MoveToBack(e3) // should be no-op

		e2 = l.InsertBefore(2, e1) // insert before front
		l.Remove(e2)
		e2 = l.InsertBefore(2, e4) // insert before middle
		l.Remove(e2)
		e2 = l.InsertBefore(2, e3) // insert before back
		l.Remove(e2)

		e2 = l.InsertAfter(2, e1) // insert after front
		l.Remove(e2)
		e2 = l.InsertAfter(2, e4) // insert after middle
		l.Remove(e2)
		e2 = l.InsertAfter(2, e3) // insert after back
		l.Remove(e2)

		// Check standard iteration.
		sum := 0
		for e := l.Front(); e != nil; e = e.Next() {
			if i, ok := e.Value.(int); ok {
				sum += i
			}
		}

		// Clear all elements by iterating
		var next *list.Element
		for e := l.Front(); e != nil; e = next {
			next = e.Next()
			l.Remove(e)
		}
	}
}

func BenchmarkExtending(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l1 := list.New()
		l2 := list.New()

		l1.PushBack(1)
		l1.PushBack(2)
		l1.PushBack(3)

		l2.PushBack(4)
		l2.PushBack(5)

		l3 := list.New()
		l3.PushBackList(l1)
		l3.PushBackList(l2)

		l3 = list.New()
		l3.PushFrontList(l2)
		l3.PushFrontList(l1)

		l3 = list.New()
		l3.PushBackList(l1)
		l3.PushBackList(l3)

		l3 = list.New()
		l3.PushFrontList(l1)
		l3.PushFrontList(l3)

		l3 = list.New()
		l1.PushBackList(l3)
		l1.PushFrontList(l3)
	}
}

func BenchmarkRemove(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := list.New()
		l.PushBack(1)
		l.PushBack(2)
		e := l.Front()
		l.Remove(e)
		l.Remove(e)
	}
}

func BenchmarkIssue4103(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l1 := list.New()
		l1.PushBack(1)
		l1.PushBack(2)

		l2 := list.New()
		l2.PushBack(3)
		l2.PushBack(4)

		e := l1.Front()
		l2.Remove(e) // l2 should not change because e is not an element of l2

		l1.InsertBefore(8, e)
	}
}

func BenchmarkIssue6349(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := list.New()
		l.PushBack(1)
		l.PushBack(2)

		e := l.Front()
		l.Remove(e)
	}
}

func BenchmarkMove(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		l := list.New()
		e1 := l.PushBack(1)
		e2 := l.PushBack(2)
		e3 := l.PushBack(3)
		e4 := l.PushBack(4)

		l.MoveAfter(e3, e3)
		l.MoveBefore(e2, e2)

		l.MoveAfter(e3, e2)
		l.MoveBefore(e2, e3)

		l.MoveBefore(e2, e4)
		e2, e3 = e3, e2

		l.MoveBefore(e4, e1)
		e1, e2, e3, e4 = e4, e1, e2, e3

		l.MoveAfter(e4, e1)
		e2, e3, e4 = e4, e2, e3

		l.MoveAfter(e2, e3)
	}
}

// Test PushFront, PushBack, PushFrontList, PushBackList with uninitialized List
func BenchmarkZeroList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var l1 = list.New()
		l1.PushFront(1)

		var l2 = list.New()
		l2.PushBack(1)

		var l3 = list.New()
		l3.PushFrontList(l1)

		var l4 = list.New()
		l4.PushBackList(l2)
	}
}

// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
func BenchmarkInsertBeforeUnknownMark(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var l list.List
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)
		l.InsertBefore(1, new(list.Element))
	}
}

// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
func BenchmarkInsertAfterUnknownMark(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var l list.List
		l.PushBack(1)
		l.PushBack(2)
		l.PushBack(3)
		l.InsertAfter(1, new(list.Element))
	}
}

// Test that a list l is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of l.
func BenchmarkMoveUnknownMark(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var l1 list.List
		e1 := l1.PushBack(1)

		var l2 list.List
		e2 := l2.PushBack(2)

		l1.MoveAfter(e1, e2)

		l1.MoveBefore(e1, e2)
	}
}
