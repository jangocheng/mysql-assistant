package dev

import (
	"container/list"

	"github.com/gin-gonic/gin"
)

// golang container
//https://studygolang.com/articles/9539

// list 是链表操作， 不是队列， Push 会加元素，  减元素需要Remove
func LearnList(c *gin.Context) {
	list := list.New()

	list.PushBack(1)
}

type Student struct {
	name  string
	score int
}

type StudentHeap []Student

func (h StudentHeap) Len() int { return len(h) }

func (h StudentHeap) Less(i, j int) bool {
	return h[i].score < h[j].score //最小堆
	//return stu[i].score > stu[j].score //最大堆
}

func (h StudentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StudentHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Student))
}

func (h *StudentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LearnLeap(c *gin.Context) {

}
