package utils

import (
	"fmt"
	"testing"
)

type MyStruct struct {
	Score int
	Value string
}

func Less(a, b MyStruct) bool {
	return a.Score < b.Score
}

func TestHeap(t *testing.T) {
	heap := New(Less)
	heap.Push(MyStruct{Score: 30, Value: "c"})
	heap.Push(MyStruct{Score: 10, Value: "a"})
	heap.Push(MyStruct{Score: 20, Value: "b"})

	item := heap.Pop()
	fmt.Println(item)
	if item.Score != 10 {
		t.Fatalf("score should be %d", item.Score)
	}
}
