package SANgo

import (lib "SANgo/internal")

// "component_filename"
func (p *Process) Send(data interface{}, dst string, timeEnd float64) interface{} {
	p._sendPacket(data, dst, SyncNetworkEvent, timeEnd)
	return p.addSync()
}


func (p *Process) DetachedSend(data interface{}, dst string, timeEnd float64) interface{} {
	return p._sendPacket(data, dst, AsyncNetworkEvent, timeEnd)
}


func (p *Process) _sendPacket(data interface{}, dst string, eventType EventType, timeEnd float64) interface{} {

	event := NewEvent(timeEnd, process, eventType, data)
	event.dst = dst

	link.Put(event, &process.env.globalQueue)
	return nil
}
