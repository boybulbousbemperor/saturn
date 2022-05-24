package tasks

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

const (
	PendingState int32 = iota
	Scheduled
	Completed
	Running
	Failed
)

type State int

type Task struct {
	ID            uuid.UUID
	Name          string
	State         int32
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskedJobEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
