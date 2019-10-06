package internal

type (
	NetworkEvent struct {
		*BaseEvent
		Dst string
	}
)


func NewNetworkEvent(timeEnd float64, dst string) *NetworkEvent{
	return &NetworkEvent{
		Dst:dst,
		BaseEvent: NewBaseEvent(timeEnd),
	}
}

func (sne *SyncNetworkEvent) GetNextWorkers() [] {

}