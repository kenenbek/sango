package SANgo

import (
	lib "SANgo/internal"
)

type Process struct {
	Pid        uint64
	Name       string
	Data       interface{}
	Resource   *Host
	resumeChan chan interface{}
	stepEnd    chan interface{}
}

func (p *Process) addSync() interface{} {
	p.stepEnd <- struct{}{}
	return <-p.resumeChan
}

func (env *environment) NewProcess(name string, host *Host) *Process {
	return &Process{
		Name:       name,
		Resource:   host,
		resumeChan: make(chan interface{}),
		stepEnd:    env.stepEnd,
	}
}
