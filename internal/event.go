package internal

import "container/heap"

type (
	Event interface {
		GetTimeEnd() int
		GetIndex() int
		GetData() interface{}
		ExecuteCallbacks()
		SetIndex(int)
	}

	BaseEvent struct {
		timeEnd float64
		status interface{}
		data interface{}
		index   uint64
		id uint64
	}
)

func (be *BaseEvent) populateNextWorker() {
	
}

func NewBaseEvent(timeEnd float64) *BaseEvent{
	return &BaseEvent{
		timeEnd:timeEnd,
	}
}

/*
=================================================================================================
Global event queue
=================================================================================================
*/

var _ heap.Interface = (*GlobalEventQueue)(nil)

type GlobalEventQueue []Event

func (eq GlobalEventQueue) Len() int { return len(eq) }

func (eq GlobalEventQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return eq[i].GetTimeEnd() < eq[j].GetTimeEnd()
}

func (eq GlobalEventQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
	eq[i].SetIndex(i) // eq[i].index = i
	eq[j].SetIndex(j) // eq[j].index = j
}

func (eq *GlobalEventQueue) Push(x interface{}) {
	n := len(*eq)
	item := x.(Event)
	item.SetIndex(n)
	*eq = append(*eq, item)
}


func (eq *GlobalEventQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	item := old[n-1]
	item.SetIndex(-1) // for safety
	*eq = old[0 : n-1]
	return item
}

