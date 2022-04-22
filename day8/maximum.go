package day8

import (
	"container/heap"
)

type ss struct {
	v []int
}

func (s *ss) Len() int           { return len(s.v) }
func (s *ss) Less(i, j int) bool { return s.v[i] < s.v[j] }
func (s *ss) Swap(i, j int)      { s.v[i], s.v[j] = s.v[j], s.v[i] }
func (s *ss) Push(x interface{}) { s.v = append(s.v, x.(int)) }

func (s *ss) Pop() interface{} {
	last := len(s.v) - 1
	if last >= 0 {
		item := s.v[last]
		s.v = s.v[0:last]
		return item
	} else {
		return nil
	}
}

func maximumProduct(nums []int, k int) int {
	ll := &ss{nums}
	heap.Init(ll)
	for ; k > 0; k-- {
		n := heap.Pop(ll).(int) + 1
		heap.Push(ll, n)
	}
	max := 1
	for i := len(ll.v) - 1; i >= 0; i-- {
		max = (max * ll.v[i]) % 1000000007
	}
	return max
}
