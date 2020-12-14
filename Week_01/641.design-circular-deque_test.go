package leetcode

import (
	"strconv"
	"testing"
)

func TestMyCircularDeque(t *testing.T) {
	circularDeque := MyCircularDequeConstructor(3)
	if ans := circularDeque.InsertLast(1); !ans {
		t.Errorf("InsertLast(1) should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.InsertLast(2); !ans {
		t.Errorf("InsertLast(2) should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.InsertFront(3); !ans {
		t.Errorf("InsertFront(3) should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.InsertFront(4); ans {
		t.Errorf("InsertFront(4) should be false but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.GetRear(); ans != 2 {
		t.Errorf("GetRear() should be 2 but got %s", strconv.Itoa(ans))
	}
	if ans := circularDeque.IsFull(); !ans {
		t.Errorf("IsFull()  should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.DeleteLast(); !ans {
		t.Errorf("DeleteLast()  should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.InsertFront(4); !ans {
		t.Errorf("InsertFront(4)  should be true but got %s", strconv.FormatBool(ans))
	}
	if ans := circularDeque.GetFront(); ans != 4 {
		t.Errorf("GetFront()  should be 4 but got %s", strconv.Itoa(4))
	}
}
