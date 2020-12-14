package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := BuildCyclelinkedList([]int{1, 2, 4}, -1)
	l2 := BuildCyclelinkedList([]int{1, 3, 4}, -1)
	if ans := fmt.Sprintf("%+v", WalkThroughList(mergeTwoLists(l1, l2))); ans != "[1 1 2 3 4 4]" {
		t.Errorf("[1, 2, 4] [1, 3, 4] should be [1 1 2 3 4 4] but got %s", ans)
	}
}
