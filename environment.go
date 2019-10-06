package SANgo

import (
	"SANgo/internal"
	"container/heap"
	"fmt"
)

type (
	environment struct {
	globalQueue     internal.GlobalEventQueue
	currentTime     float64
	waiters         map[internal.Event]*Process
	workers         map[uint64]*Process
	workerListeners map[string]*Process
	nextWorkers     []*Process
	daemonList      map[uint64]*Process
	shouldStop      bool
	stepEnd         chan interface{}
	systemResources []Closeable
}
	Closeable interface {
		Close ()
	}
)


func NewEnvironment() *environment {
	queue := make(internal.GlobalEventQueue, 0)
	heap.Init(&queue)

	e := &environment{
		globalQueue:     queue,
		workers:         make(map[uint64]*Process),
		daemonList:      make(map[uint64]*Process),
		workerListeners: make(map[string]*Process),
		nextWorkers:     make([]*Process, 0),
		stepEnd:         make(chan interface{}),
	}

	return e
}


func (env *environment) populateNextWorkers(event internal.Event) {

	switch event {
	case *internal.NetworkEvent:
		ne := event.(*internal.NetworkEvent)
		lp, ok := env.workerListeners[ne.Dst]
		if !ok {
			panic(fmt.Sprintf("No such listener"))
		}
		env.nextWorkers = append(env.nextWorkers, lp)
	default:
		p, ok := env.waiters[event]
		if ok {
			env.nextWorkers = append(env.nextWorkers, p)
		}
	}
}