package model

import "encoding/gob"

// EventType
const (
	EventError = iota - 2
	EventNone
	_
	EventSync
	EventNodeConfChanged
	EventNodeRegister
	EventNodeStatus
	EventNodeProc
	EventNodeExit
	EventCmdStart
	EventCmdStop
	EventCmdRestart
	EventCollectProc
	EventCollectPProf
	EventCollectGorpc
	EventCollectRedisPool
	EventCollectProfile

	EventOperate
	EventOperateBatch
	EventOperateRollback
)

func init() {
	gob.RegisterName(PkgPrefix + "Event", Event{})
}

// Event .
type Event struct {
	EventType int
	Data      interface{}
}

// NewEvent return EventTypeNone event
func NewEvent() Event {
	return Event{EventType: EventNone}
}
