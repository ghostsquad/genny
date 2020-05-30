// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/ghostsquad/genny

package queue

// Float32Queue is a queue of Float32s.
type Float32Queue struct {
	items []float32
}

func NewFloat32Queue() *Float32Queue {
	return &Float32Queue{items: make([]float32, 0)}
}
func (q *Float32Queue) Push(item float32) {
	q.items = append(q.items, item)
}
func (q *Float32Queue) Pop() float32 {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
